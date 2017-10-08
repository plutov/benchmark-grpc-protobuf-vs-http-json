package benchmarks

import (
	"github.com/plutov/benchmark-grpc-rest/grpc"
	"github.com/plutov/benchmark-grpc-rest/grpc/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
	"testing"
	"time"
)

func init() {
	go grpc.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPC(b *testing.B) {
	conn, _ := g.Dial("127.0.0.1:60000", g.WithInsecure())
	client := proto.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		client.CreateUser(context.Background(), &proto.Request{
			Email:    "foo@bar.com",
			Name:     "Bench",
			Password: "bench",
		})
	}
}
