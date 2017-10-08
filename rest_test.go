package benchmarks

import (
	"bytes"
	//"encoding/json"
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
	buf := bytes.NewBufferString(`{"email":"foo@bar.com","name":"Bench","password":"bench"}`)

	for n := 0; n < b.N; n++ {
		client.Post("http://127.0.0.1:60001/", "application/json", buf)
	}
}
