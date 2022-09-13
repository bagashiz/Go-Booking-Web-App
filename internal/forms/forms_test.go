package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/foo", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/foo", nil)
	form := New(r.PostForm)

	form.Required("foo", "bar", "buzz")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/foo", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields are present")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("foo")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.MinLength("foo", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("foo")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some_value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minLength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another field", "abc123")
	form = New(postedValues)

	form.MinLength("another field", 1)
	if !form.Valid() {
		t.Error("shows minLength of  1 is not met when it is")
	}

	isError = form.Errors.Get("another fiels")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("foo")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "foo@bar.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when it should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "foo")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got a valid email for invalid email address")
	}
}
