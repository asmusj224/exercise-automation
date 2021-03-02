package services

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockHttpClient struct{}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	json := `{"sid":"12345"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

func TestSendSMS(t *testing.T) {
	mockClient := &MockHttpClient{}
	service := NewMessageService(mockClient)
	msg := "Hi"
	res, err := service.SendSMS("+12532034540", msg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res, "12345")
}
