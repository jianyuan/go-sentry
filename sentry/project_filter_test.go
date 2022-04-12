package sentry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectFilterService_GetWithLegacyExtension(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/powerful-abolitionist/filters/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, getWithLegacyExtensionHeader)
	})

	client := NewClient(httpClient, nil, "")
	filterConfig, _, err := client.ProjectFilter.GetFilterConfig("the-interstellar-jurisdiction", "powerful-abolitionist")
	assert.NoError(t, err)

	expected := FilterConfig{
		LegacyBrowsers:   []string{"ie_pre_9"},
		BrowserExtension: false,
	}
	assert.Equal(t, &expected, filterConfig)
}

func TestProjectFilterService_GetWithoutLegacyExtension(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/powerful-abolitionist/filters/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, getWithoutLegacyExtensionHeader)
	})

	client := NewClient(httpClient, nil, "")
	filterConfig, _, err := client.ProjectFilter.GetFilterConfig("the-interstellar-jurisdiction", "powerful-abolitionist")
	assert.NoError(t, err)

	expected := FilterConfig{
		LegacyBrowsers:   nil,
		BrowserExtension: true,
	}
	assert.Equal(t, &expected, filterConfig)
}

func readRequestBody(r *http.Request) string {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		panic(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")
	return str
}

func TestBrowserExtensionFilter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/test_org/test_project/filters/browser-extensions/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		body := readRequestBody(r)
		assert.Equal(t, body, `{"active":true}`)
		w.Header().Set("Content-Type", "application/json")
	})
	client := NewClient(httpClient, nil, "")
	_, err := client.ProjectFilter.UpdateBrowserExtensions("test_org", "test_project", true)
	assert.NoError(t, err)
}

func TestLegacyBrowserFilter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/test_org/test_project/filters/legacy-browsers/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		body := readRequestBody(r)
		assert.Equal(t, body, `{"subfilters":["ie_pre_9","ie10"]}`)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
	client := NewClient(httpClient, nil, "")
	browsers := []string{"ie_pre_9", "ie10"}
	_, err := client.ProjectFilter.UpdateLegacyBrowser("test_org", "test_project", browsers)
	assert.NoError(t, err)
}

var (
	getWithLegacyExtensionHeader = `[
		{
			"id":"browser-extensions",
			"active":false,
			"description":"description_1",
			"name":"name_1",
			"hello":"hello_1"
		},
		{
			"id":"localhost",
			"active":false,
			"description":"description_2",
			"name":"name_2",
			"hello":"hello_2"
		},
		{
			"id":"legacy-browsers",
			"active":["ie_pre_9"],
			"description":"description_3",
			"name":"name_3",
			"hello":"hello_3"
		},
		{
			"id":"web-crawlers",
			"active":true,
			"description":"description_4",
			"name":"name_4",
			"hello":"hello_4"
		}
	]`
	getWithoutLegacyExtensionHeader = `[
		{
			"id":"browser-extensions",
			"active":true,
			"description":"description_1",
			"name":"name_1",
			"hello":"hello_1"
		},
		{
			"id":"localhost",
			"active":false,
			"description":"description_2",
			"name":"name_2",
			"hello":"hello_2"
		},
		{
			"id":"legacy-browsers",
			"active":false,
			"description":"description_3",
			"name":"name_3",
			"hello":"hello_3"
		},
		{
			"id":"web-crawlers",
			"active":true,
			"description":"description_4",
			"name":"name_4",
			"hello":"hello_4"
		}
	]`
)
