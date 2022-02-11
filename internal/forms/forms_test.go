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

	if !form.MinLength("a", 1) {
		t.Error("form shows that MinLength isn't correct when it is")
	}

	if form.MinLength("a", 3) {
		t.Error("form shows that MinLength correct when it isn't")
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
