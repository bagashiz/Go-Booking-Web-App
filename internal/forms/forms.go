package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

// Required is a function that checks for required fields specified in the parameter.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

// Has is a function that checks if form field is in POST and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	return x != ""
}

// Valid is a function that returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// MinLength is a function that checks for minimum character length of a field.
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail is a function that checks for valid email address.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address.")
	}
}
