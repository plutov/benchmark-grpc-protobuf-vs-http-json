package benchmarks

import (
	"testing"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf"
	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, _ := g.Dial("127.0.0.1:60000", g.WithInsecure())
	client := proto.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client proto.APIClient, b *testing.B) {
	_, err := client.CreateUser(context.Background(), &proto.Request{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	})
	if err != nil {
		b.Fatal(err.Error())
	}
}
