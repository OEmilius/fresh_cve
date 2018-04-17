package main

import (
	"log"
	"time"

	pb "github.com/OEmilius/fresh_cve/cve"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOffenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//for get all cve need to set Filter = all
	r, err := c.GetCves(ctx, &pb.CveRequest{Filter: "all"})
	if err != nil {
		log.Fatalf("no answer from server: %v", err)
	}
	log.Printf("all cves is: %s", r.Cves)
}
