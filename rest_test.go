package benchmarks

import (
	"bytes"
	"github.com/plutov/benchmark-grpc-rest/rest"
	"net/http"
	"testing"
	"time"
)

func init() {
	rest.Start()
	time.Sleep(time.Second)
}

func BenchmarkREST(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		buf := bytes.NewBufferString(`{"email":"foo@bar.com","name":"Bench","password":"bench"}`)
		client.Post("http://127.0.0.1:60001/", "application/json", buf)
	}
}
