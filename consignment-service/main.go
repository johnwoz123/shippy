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
