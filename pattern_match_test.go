package pathmatch_test

import (
	"github.com/reiver/go-pathmatch"

	"testing"
)

func TestPatternMatch(t *testing.T) {
	tests := []struct{
		Pattern  string
		Path     string
		Expected bool
	}{
		{
			Pattern: "/v1/help",
			Path:    "/v1/help",
			Expected: true,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/help/",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/help/me",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/help/me/",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/helping",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/helping/",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v2/help",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v2/HELP",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/apple",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/banana",
			Expected: false,
		},
		{
			Pattern: "/v1/help",
			Path:    "/v1/cherry",
			Expected: false,
		},



		{
			Pattern: "/v1/help/",
			Path:    "/v1/help/",
			Expected: true,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/help",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/help/me",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/help/me/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/helping",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/helping/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v2/help/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v2/HELP/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/apple/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/banana/",
			Expected: false,
		},
		{
			Pattern: "/v1/help/",
			Path:    "/v1/cherry/",
			Expected: false,
		},



		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1/user/123",
			Expected: true,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1/user/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1/user/123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "//v1/user/123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1//user/123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1/user//123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "//v1//user/123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "/v1//user//123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "//v1//user//123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}",
			Path:    "//v1//user//123//",
			Expected: false,
		},



		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1/user/123/",
			Expected: true,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1/user/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1/user/123",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "//v1/user/123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1//user/123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1/user//123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "//v1//user/123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "/v1//user//123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "//v1//user//123/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/",
			Path:    "//v1//user//123//",
			Expected: false,
		},



		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}",
			Path:    "/v1/user/123/contact/e-mail",
			Expected: true,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}",
			Path:    "/v1/user/123/contact/e-mail/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}",
			Path:    "/v2/user/123/contact/e-mail",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}",
			Path:    "/v2/user/123/contact/e-mail",
			Expected: false,
		},



		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}/",
			Path:    "/v1/user/123/contact/e-mail/",
			Expected: true,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}/",
			Path:    "/v1/user/123/contact/e-mail",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}/",
			Path:    "/v2/user/123/contact/e-mail/",
			Expected: false,
		},
		{
			Pattern: "/v1/user/{user_id}/contact/{contact_type}/",
			Path:    "/v2/user/123/contact/e-mail/",
			Expected: false,
		},



		{
			Pattern: "/v1/company/{company_name}",
			Path:    "/v1/company/acme",
			Expected: true,
		},
		{
			Pattern: "/v1/company/{company_name}",
			Path:    "/v1/company/acme/",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}",
			Path:    "/v2/company/acme",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}",
			Path:    "/v1/user/acme",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}",
			Path:    "/v1/COMPANY/acme",
			Expected: false,
		},



		{
			Pattern: "/v1/company/{company_name}/",
			Path:    "/v1/company/acme/",
			Expected: true,
		},
		{
			Pattern: "/v1/company/{company_name}/",
			Path:    "/v1/company/acme",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}/",
			Path:    "/v2/company/acme/",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}/",
			Path:    "/v1/user/acme/",
			Expected: false,
		},
		{
			Pattern: "/v1/company/{company_name}/",
			Path:    "/v1/COMPANY/acme/",
			Expected: false,
		},



		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.jsonld",
			Expected: true,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.jsonld",
			Expected: true,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.jsonl",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.jsonl",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.json",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.json",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.jso",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.jso",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.js",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.js",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.j",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.j",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c.",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ed7ba470-8e54-465e-825c-99712043e01c",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/object/ED7BA470-8E54-465E-825C-99712043E01C",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/OBJECT/ed7ba470-8e54-465e-825c-99712043e01c.jsonld",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/-/OBJECT/ED7BA470-8E54-465E-825C-99712043E01C.jsonld",
			Expected: false,
		},
		{
			Pattern: "/-/object/{uuid}.jsonld",
			Path:    "/apple/banana/cherry/ed7ba470-8e54-465e-825c-99712043e01c.jsonld",
			Expected: false,
		},



		{
			Pattern: "/-/outbox({begin}).jsonld",
			Path:    "/-/outbox(0).jsonld",
			Expected: true,
		},
		{
			Pattern: "/-/outbox({begin}).jsonld",
			Path:    "/-/outbox(10).jsonld",
			Expected: true,
		},

		{
			Pattern: "/-/outbox({begin},{end}).jsonld",
			Path:    "/-/outbox(0,10).jsonld",
			Expected: true,
		},
		{
			Pattern: "/-/outbox({begin},{end}).jsonld",
			Path:    "/-/outbox(10,20).jsonld",
			Expected: true,
		},
	}

	for testNumber, test := range tests {
		var pattern pathmatch.Pattern

		if err := pathmatch.CompileTo(&pattern, test.Pattern); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\t: PATTERN:  %q", test.Pattern)
			t.Errorf("\t: PATH:     %q", test.Path)
			t.Errorf("\t: EXPECTED: %t", test.Expected)
			continue
		}

		matched, err := pattern.Match(test.Path)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\t: PATTERN:  %q", test.Pattern)
			t.Errorf("\t: PATH:     %q", test.Path)
			t.Errorf("\t: EXPECTED: %t", test.Expected)
			continue
		}

		if expected, actual := test.Expected, matched; expected != actual {
			t.Errorf("For test #%d, expected %t, but actually got %t.", testNumber, expected, actual)
			t.Errorf("\t: PATTERN:  %q", test.Pattern)
			t.Errorf("\t: PATH:     %q", test.Path)
			t.Errorf("\t: EXPECTED: %t", test.Expected)
			continue
		}
	}
}
