//go:generate protoc -I ../cve --go_out=plugins=grpc:../cve ../cve/cve.proto

package grpc_server

import (
	"log"
	"net"

	cache "github.com/OEmilius/fresh_cve/cache"
	pb "github.com/OEmilius/fresh_cve/cve"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ListenPort string = ":50051"

var CACHE *cache.Cache

type server struct{}

func (s *server) GetCves(ctx context.Context, in *pb.CveRequest) (*pb.CveReply, error) {
	cves := []*pb.Cve{}
	if in.Filter == "all" {
		cves = CACHE.GetAllCve_forGrpc()
		return &pb.CveReply{cves}, nil
	}
	return &pb.CveReply{cves}, nil
}

func Start() {
	lis, err := net.Listen("tcp", ListenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server listen:", ListenPort)
	s := grpc.NewServer()
	pb.RegisterOffenderServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
