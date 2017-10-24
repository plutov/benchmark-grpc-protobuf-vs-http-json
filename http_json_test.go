package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

func init() {
	go httpjson.Start()
	time.Sleep(time.Second)
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
		b.Fatalf("http request failed: %v", err)
	}
	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target httpjson.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
