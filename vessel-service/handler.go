package main

import (
	"context"
	pb "github.com/AzusaChino/daphne/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, specification *pb.Specification, response *pb.Response) error {
	vessel, err := h.repository.FindAvailable(ctx, MarshalSpecification(specification))
	if err != nil {
		return err
	}
	response.Vessel = UnmarshalVessel(vessel)
	return nil
}

func (h *handler) Create(ctx context.Context, vessel *pb.Vessel, response *pb.Response) error {
	if err := h.repository.Create(ctx, MarshalVessel(vessel)); err != nil {
		return err
	}
	response.Vessel = vessel
	return nil
}
