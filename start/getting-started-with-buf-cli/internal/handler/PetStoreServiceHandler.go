package handler

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"context"
)

type PetStoreServiceHandler struct {
}

func (p PetStoreServiceHandler) GetPet(ctx context.Context, request *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	//TODO implement me
	panic("implement me")
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
