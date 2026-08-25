package api

import (
	"strings"
	"testing"
)

func TestXmlMarshalWithSelfClosingTags(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"alphanumeric", "abc123"},
		{"at and hash", "p@ss#word"},
		{"exclamation and percent", "p!ss%word"},
		{"colon semicolon comma", "p:ss;wo,rd"},
		{"brackets and parens", "p[a](b)word"},
		{"quote requiring escape", `p"ss'word`},
		{"ampersand requiring escape", "p&ss&word"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := AaaLoginRequest{
				InName:     "user",
				InPassword: tt.password,
			}
			data, err := xmlMarshalWithSelfClosingTags(req)
			if err != nil {
				t.Fatalf("xmlMarshalWithSelfClosingTags: %v", err)
			}
			got := string(data)
			if !strings.HasSuffix(got, "/>") {
				t.Errorf("expected self-closing tag, got: %s", got)
			}
			if strings.Contains(got, "></aaaLogin>") {
				t.Errorf("expected element to be collapsed to self-closing form, got: %s", got)
			}
			if strings.Contains(got, "&#34;") || strings.Contains(got, "&#39;") {
				t.Errorf("expected quote/apostrophe as named entities, got numeric character reference: %s", got)
			}
		})
	}
}
