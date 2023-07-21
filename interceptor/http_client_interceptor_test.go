package interceptor_test

import (
	"github.com/mateusmrangel/client-encryption-go/interceptor"
	"github.com/mateusmrangel/client-encryption-go/jwe"
	"testing"
)

func TestHttpClientInterceptor(t *testing.T) {
	jwe := jwe.JWEConfig{}
	client, e := interceptor.GetHttpClient(jwe, nil)
	if e != nil || client == nil {
		t.Errorf("Expected http.Client, but got %v", e)
	}
}
