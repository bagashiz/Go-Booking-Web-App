package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", nil, http.StatusOK},
	{"about", "/about", "GET", nil, http.StatusOK},
	{"executive", "/executive", "GET", nil, http.StatusOK},
	{"deluxe", "/deluxe", "GET", nil, http.StatusOK},
	{"contact", "/contact", "GET", nil, http.StatusOK},
	{"search-availability", "/search-availability", "GET", nil, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", nil, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "checkInDate", value: "2020-01-01"},
		{key: "checkOutDate", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "checkInDate", value: "2020-01-01"},
		{key: "checkOutDate", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Doe"},
		{key: "email", value: "john@doe.com"},
		{key: "phone", value: "111222333"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else { // POST
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
