package grpcservice

import (
	"github.com/tanphamhaiduong/go/delta/internal/modules"
	"github.com/tanphamhaiduong/go/delta/pb"
	"google.golang.org/grpc"
)

// DeltaServiceImpl ...
type DeltaServiceImpl struct {
	handlers modules.Handler
}

// NewGrpcService ...
func NewGrpcService(handlers modules.Handler) *DeltaServiceImpl {
	return &DeltaServiceImpl{
		handlers: handlers,
	}
}

// Register ...
func (d *DeltaServiceImpl) Register(server *grpc.Server) {
	pb.RegisterDeltaServiceServer(server, d)
}
