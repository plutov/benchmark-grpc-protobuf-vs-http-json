package benchmarks

import (
	"bytes"
	"net/http"
	"testing"
)

func BenchmarkREST(b *testing.B) {
	client := &http.Client{}
	buf := bytes.NewBufferString(`{"email":"foo@bar.com","name":"Bench","password":"bench"}`)

	for n := 0; n < b.N; n++ {
		client.Post("http://127.0.0.1:60001/", "application/json", buf)
	}
}
