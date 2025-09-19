package auth

import (
	"net/http"
	"net/textproto"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		http.CanonicalHeaderKey("Content-Type"):           {"application/json"},
		textproto.CanonicalMIMEHeaderKey("Authorization"): {"ApiKey TESTTOKEN"},
	}
	got, _ := GetAPIKey(headers)
	want := "TEST"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
