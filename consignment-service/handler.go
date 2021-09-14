package main

import (
	"context"
	pb "github.com/AzusaChino/daphne/consignment-service/proto/consignment"
	vesselPb "github.com/AzusaChino/daphne/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselPb.VesselService
}

func (h handler) CreateConsignment(ctx context.Context, consignment *pb.Consignment, response *pb.Response) error {
	panic("implement me")
}
