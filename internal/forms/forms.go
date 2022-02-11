package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Creates custom form struck embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for list of fields not blank
func (f *Form) Required(field ...string) {
	for _, field := range field {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//Has checks if form contains field and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	return x != ""
}

//MinLength check for minimum length of the field
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}

//IsEmail checks for valid email address
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "invalid email address")
		return false
	}

	return true
}
