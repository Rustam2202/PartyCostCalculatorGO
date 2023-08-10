package main

import (
	"context"
	"party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonHandler struct {
	service service.PersonService
}

func (h *PersonHandler) AddPerson(ctx context.Context, pb *proto.PersonCreate) (*proto.Id, error) {
	id, err := h.service.NewPerson(ctx, pb.Name)
	if err != nil {
		return nil, err
	}
	return &proto.Id{Id: id}, nil
}

func (h *PersonHandler) GetPerson(ctx context.Context, pb *proto.Id) (*proto.Person, error) {
	person, err := h.service.GetPersonById(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	result := &proto.Person{
		Id:     person.Id,
		Name:   person.Name,
		Events: []*proto.Event{},
	}
	for _, ev := range person.Events {
		result.Events = append(result.Events, &proto.Event{
			Id:   ev.Id,
			Name: ev.Name,
			Date: ev.Date.Format("2006-01-02"),
		})
	}
	return result, nil
}

func (h *PersonHandler) UpdatePerson(ctx context.Context, pb *proto.PersonUpdate) (*proto.Empty, error) {
	err := h.service.UpdatePerson(ctx, pb.Id, pb.Name)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *PersonHandler) DeletePerson(ctx context.Context, pb *proto.Id) (*proto.Empty, error) {
	err := h.service.DeletePersonById(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
