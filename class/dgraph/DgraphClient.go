package DgraphClient

import (
	"context"
	"log"
	// "fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	grpc "google.golang.org/grpc"
)

func NewClient() *dgo.Dgraph {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func RunQuery(q string) *api.Response{
	dg := NewClient()
	ctx := context.Background()
	resp, errDB := dg.NewTxn().Query(ctx, q)
	if errDB != nil {
		log.Fatal(errDB)
	}
	return resp
}