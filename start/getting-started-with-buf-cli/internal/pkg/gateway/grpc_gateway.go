package gateway

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCGateway struct {
	gRPCConn *grpc.ClientConn
}

func NewGRPCGateway() *GRPCGateway {
	return &GRPCGateway{}
}

type RegisterServiceHandlerFunc func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
type RegisterServiceServerFunc func(context.Context, *grpc.ServiceRegistrar)

func (g *GRPCGateway) CreateHttpHandler(
	ctx context.Context,
	endpoint string,
	funcs ...RegisterServiceHandlerFunc,

) http.Handler {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		panic(err)
	}
	g.gRPCConn = conn

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						EmitUnpopulated: true,
						UseEnumNumbers:  true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			},
		),
	)

	for _, v := range funcs {
		_ = v(ctx, mux, conn)
	}
	return handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "ResponseType"}),
	)(mux)
}
