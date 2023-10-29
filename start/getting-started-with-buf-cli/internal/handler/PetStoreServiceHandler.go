package handler

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PetStoreServiceHandler struct {
}

func (p PetStoreServiceHandler) GetPet(ctx context.Context, request *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	return &petv1.GetPetResponse{
		Pet: &petv1.Pet{
			PetType:   petv1.PetType_PET_TYPE_CAT,
			PetId:     "1",
			Name:      "name-1",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

func (p PetStoreServiceHandler) PutPet(ctx context.Context, request *petv1.PutPetRequest) (*petv1.PutPetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetStoreServiceHandler) DeletePet(ctx context.Context, request *petv1.DeletePetRequest) (*petv1.DeletePetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewPetStoreServiceHandler() petv1.PetStoreServiceServer {
	return &PetStoreServiceHandler{}
}
