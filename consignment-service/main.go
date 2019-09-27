package main

import (
	"sync"

	pb "github.com/johnwoz123/shippy-service-consignment/proto/consignment"
)

const (
	port = "50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	myMutex      sync.RWMutex
	consignments []*pb.Consignment
}
