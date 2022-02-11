package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/123", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows invalid when it valid")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	if form.Has("a") {
		t.Error("form shows that has property when it isn't")
	}

	postedData.Add("a", "a")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("form shows that hasn't property when it is")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "ab")

	form := New(postedData)

	form.MinLength("a", 1)
	isError := form.Errors.Get("a")
	if !form.Valid() {
		t.Error("form shows that MinLength isn't correct when it is")
	}

	if isError != "" {
		t.Error("form shows that MinLength isn't has error")
	}

	form.MinLength("a", 3)
	isError = form.Errors.Get("a")
	if form.Valid() {
		t.Error("form shows that MinLength correct when it isn't")
	}

	if isError == "" {
		t.Error("form shows that MinLength is has error")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "ts@test.com")

	form := New(postedData)

	if !form.IsEmail("email") {
		t.Error("form shows that field is not email whe it is")
	}

	postedData.Set("email", "ts1test.com")
	form = New(postedData)

	if form.IsEmail("email") {
		t.Error("form shows that field is email whe it is not")
	}
}
