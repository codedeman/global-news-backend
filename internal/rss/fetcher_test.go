package rss

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch_Success(t *testing.T) {
	expected := []byte("hello rss feed")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(expected)
	}))
	defer server.Close()

	data, err := Fetch(server.URL)

	require.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestFetch_InvalidURL(t *testing.T) {
	_, err := Fetch("http://127.0.0.1:0/invalid")

	assert.Error(t, err)
}
