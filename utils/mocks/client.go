package mocks

import "net/http"

type MockClient struct {
	DoFunc GetDoFunc
}

type GetDoFunc func(req *http.Request) (*http.Response, error)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
