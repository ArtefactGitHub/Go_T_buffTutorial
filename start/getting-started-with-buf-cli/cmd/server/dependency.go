package server

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"Go_T_buffTutorial/internal/pkg/gateway"
)

func RegisterServiceHandlerFuncs() []gateway.RegisterServiceHandlerFunc {
	return []gateway.RegisterServiceHandlerFunc{
		petv1.RegisterPetStoreServiceHandler,
	}
}

//func RegisterServiceServerFuncs() []gateway.RegisterServiceServerFunc {
//	return []gateway.RegisterServiceServerFunc{
//		func(ctx context.Context, registrar *grpc.ServiceRegistrar) {
//			petv1.RegisterPetStoreServiceServer(*registrar, handler.NewPetStoreServiceHandler())
//		},
//	}
//}
