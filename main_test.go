package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dhruv42/mocking/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	jsonResponse := `[{
		"full_name": "mock-repo"
	}]`

	postBody := map[string]string{
		"name": "morpheus",
		"job":  "leader",
	}

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &mocks.MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	resp, err := CreateUser(postBody)
	fmt.Println(resp)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
