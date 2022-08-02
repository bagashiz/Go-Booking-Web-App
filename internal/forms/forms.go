package forms

import (
	"net/http"
	"net/url"
)

// Form is a custom form struct that embeds a url.Values and errors struct.
type Form struct {
	url.Values
	Errors errors
}

// New is a function that initializes a form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has is a function that checks if form field is in POST and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}

// Valid is a function that returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
