package forms

import (
		"fmt"
		"net/url"
		"regexp"
		"strings"
		"unicode/utf8"
)
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Form struct {
		url.Values
		Errors errors
}

func (f *Form) MinLength(field string, d int) {
		value := f.Get(field)
		if value == "" {
				f.Errors.Add(field, "This field can't be empty")
				return
		}
		if utf8.RuneCountInString(value) < d {
				f.Errors.Add(field, fmt.Sprintf("This field should be minimum %d characters", d))
		}
}

func (f *Form) MatchPattern(field string, pattern *regexp.Regexp) {
		value := f.Get(field)
		if value == "" {
				return
		}

		if !pattern.MatchString(value) {
				f.Errors.Add(field, "This field is invalid")
		}
 }

func New(data url.Values) *Form {
		return &Form{
				data,
				errors(map[string][]string{}),
		}
}

func (f *Form) Required(fields ...string) {
		for _, field := range fields {
				value := f.Get(field)
				if strings.TrimSpace(value) == "" {
						f.Errors.Add(field, "This field can't be empty")
				}
		}
}

func (f *Form) MaxLength(field string, d int) {
		value := f.Get(field)
		if value == "" {
				return
		}
		if utf8.RuneCountInString(value) > d {
				f.Errors.Add(field, fmt.Sprintf("This field is over %d characters", d))
		}

}

func (f *Form) PermittedValues(field string, opts ...string) {
		value := f.Get(field)
		if value == "" {
				return
		}
		for _, opt := range opts {
				if value == opt {
						return
				}
		}
		f.Errors.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
		return len(f.Errors) == 0
}
