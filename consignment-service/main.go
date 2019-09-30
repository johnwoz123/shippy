package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/johnwoz123/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository fake database for now
type Repository struct {
	myMutex      sync.RWMutextex
	consignments []*pb.Consignment
}

//Create - creates a new Consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.myMutex.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.myMutex.Unlock()
	return consignment, nil
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Consignment, error) {
	// Save the consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return the correct message matching the response created in the protobuf definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}
func main() {
	repo := &Repository{}
	// set up grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %w", port)
	}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)
	log.Println("RUNNING ON PORT: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
