package benchmarks

import (
	"bytes"
	"net/http"
	"testing"
)

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

	defer resp.Body.Close()
}
