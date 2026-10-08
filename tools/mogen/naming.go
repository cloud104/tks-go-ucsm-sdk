package main

import (
	"strings"
	"unicode"
)

// initialisms são palavras que viram sigla maiúscula no nome Go
// (peerChassisId -> PeerChassisID). Acrescente conforme necessário.
var initialisms = map[string]string{
	"acl":  "ACL",
	"dns":  "DNS",
	"id":   "ID",
	"ip":   "IP",
	"uuid": "UUID",
}

// typeName converte o nome de classe XML no nome do tipo Go:
// vnicEtherIf -> VnicEtherIf.
func typeName(class string) string {
	if class == "" {
		return ""
	}
	r := []rune(class)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// fieldName converte o nome do atributo XML no nome do campo Go,
// aplicando as siglas de initialisms: peerChassisId -> PeerChassisID,
// extIPPoolName -> ExtIPPoolName.
func fieldName(attr string) string {
	var b strings.Builder
	for _, w := range splitWords(attr) {
		if up, ok := initialisms[strings.ToLower(w)]; ok {
			b.WriteString(up)
			continue
		}
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
		b.WriteString(string(r))
	}
	return b.String()
}

// splitWords quebra um identificador camelCase em palavras, tratando
// sequências de maiúsculas como sigla: "extIPPoolName" -> [ext IP Pool Name].
func splitWords(s string) []string {
	r := []rune(s)
	var words []string
	start := 0
	for i := 1; i < len(r); i++ {
		prev, cur := r[i-1], r[i]
		boundary := false
		switch {
		case unicode.IsLower(prev) && unicode.IsUpper(cur):
			boundary = true // fooBar
		case unicode.IsDigit(prev) && unicode.IsUpper(cur):
			boundary = true // ipv4If
		case unicode.IsUpper(prev) && unicode.IsUpper(cur) &&
			i+1 < len(r) && unicode.IsLower(r[i+1]):
			boundary = true // IPPool -> IP | Pool
		}
		if boundary {
			words = append(words, string(r[start:i]))
			start = i
		}
	}
	return append(words, string(r[start:]))
}
