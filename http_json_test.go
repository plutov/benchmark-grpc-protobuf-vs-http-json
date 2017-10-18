package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	ID      string `json:"id"`
}

func init() {
	go httpjson.Start()
}

func BenchmarkHTTPJSON(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doPost(client, b)
	}
}

func doPost(client *http.Client, b *testing.B) {
	buf := bytes.NewBufferString(`{"email":"foo@bar.com","name":"Bench","password":"bench"}`)

	resp, err := client.Post("http://127.0.0.1:60001/", "application/json", buf)
	if err != nil {
		b.Fatal(err.Error())
	}

	// We need to parse response to have a fair comparison as gRPC does it
	var target Response
	json.NewDecoder(resp.Body).Decode(target)
	resp.Body.Close()
}
