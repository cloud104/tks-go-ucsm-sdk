package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
)

// xmlMarshalWithSelfClosingTags post-processes results from xml.Marshal into XML
// document where empty XML elements use self-closing tags.
//
// The XML standard states that self-closing tags are permitted.
//
// https://www.w3.org/TR/REC-xml/#dt-empty
// https://www.w3.org/TR/REC-xml/#d0e2480
//
// According to the XML standard an empty element with a start and end tag is
// semantically the same as an empty element with a self-closing tag.
//
// Unfortunately not all XML parsers can understand both. Such is the
// case with Cisco UCS Manager API, which expects empty elements to use
// self-closing tags only.
//
// As of now XML marshaling in Go always uses start and end tags,
// which results in XML elements like the one below.
//
//	<Person name="me"></Person>
//
// Above XML elements cannot be parsed by the remote Cisco UCS API endpoint,
// and such API calls result in parse error returned to the client.
//
// Currently the Go team is considering implementing XML self-closing tags,
// which progress can be tracked in the issue below.
//
// https://github.com/golang/go/issues/21399
//
// Until support for XML self-closing tags (if ever) becomes real in Go
// we need to ensure compatibility with the Cisco UCS API by doing that ourselves.
//
// In a separate thread @rsc also suggested a similar approach by using
// strings.Replace(), even though such thing is not ideal and should
// hopefully one day make into the language as a feature.
//
// https://groups.google.com/forum/#!topic/golang-nuts/guG6iOCRu08
func xmlMarshalWithSelfClosingTags(in any) ([]byte, error) {
	data, err := xml.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}
	// encoding/xml escapes '"' and '\'' within attribute values using numeric
	// character references (&#34; and &#39;) rather than the equivalent named
	// entities (&quot; and &apos;). The Cisco UCS Manager XML parser only
	// recognizes the five predefined named entities and does not decode numeric
	// character references, so a password (or other attribute value) containing
	// a literal quote is received as the literal characters "&#34;" instead of
	// '"', silently corrupting the value. Rewrite to the named form, which is
	// semantically identical XML but is what the UCS parser actually understands.
	data = bytes.ReplaceAll(data, []byte("&#34;"), []byte("&quot;"))
	data = bytes.ReplaceAll(data, []byte("&#39;"), []byte("&apos;"))
	// xml.Marshal always escapes literal '<' and '>' within attribute values, so
	// any '<' or '>' in the marshaled output can only be a tag delimiter. That
	// means we don't need to whitelist which characters are allowed inside a tag
	// (which previously broke on passwords or other attribute values containing
	// characters like '@' or '#') - it's enough to match anything that isn't '<' or '>'.
	re := regexp.MustCompile(`<([^/>][^<>]*)>\s*<(/\w*)>`)
	newData := re.ReplaceAllString(string(data), "<$1/>")
	return []byte(newData), nil
}
