package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	router     Router
	testServer *httptest.Server
)

type MockRouter struct {
	RouterImpl
}

func (router MockRouter) Path() string {
	return "/"
}

func TestMain(m *testing.M) {
	router = &MockRouter{}
	testServer = httptest.NewServer(http.HandlerFunc(route(router)))

	os.Exit(m.Run())
}

func TestHandleGet(t *testing.T) {
	resp, err := http.Get(testServer.URL)

	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HandleGet: Expected response StatusCode to be: %d, got: %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestHandlePost(t *testing.T) {
	resp, err := http.Post(testServer.URL, "text/plain", strings.NewReader("test"))

	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HandlePost: Expected response StatusCode to be: %d, got: %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestHandlePut(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, testServer.URL, strings.NewReader("test"))
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HandlePut: Expected response StatusCode to be: %d, got: %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestHandleDelete(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, testServer.URL, strings.NewReader("test"))
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HandlePut: Expected response StatusCode to be: %d, got: %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestHandlePatch(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPatch, testServer.URL, strings.NewReader("test"))
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HandlePut: Expected response StatusCode to be: %d, got: %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
