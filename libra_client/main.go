package main

import (
	"context"
	"log"
	"time"

	ac "github.com/phlip9/libra_example/admission_control"
	t "github.com/phlip9/libra_example/types"
	"google.golang.org/grpc"
)

const (
	address = "ac.testnet.libra.org:8000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC AdmissionControl client
	c := ac.NewAdmissionControlClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make an empty UpdateToLatestLedger request
	m := t.UpdateToLatestLedgerRequest{ClientKnownVersion: 0, RequestedItems: []*t.RequestItem{}}
	r, err := c.UpdateToLatestLedger(ctx, &m)
	if err != nil {
		log.Fatalf("fail: %v", err)
	}

	// Print the results
	log.Printf("success: %s", r)
}
