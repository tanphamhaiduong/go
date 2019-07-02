package grpcservice

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
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

// CompanyList ...
func (d *DeltaServiceImpl) CompanyList(ctx context.Context, req *pb.CompanyListRequest) (*pb.CompanyListResponse, error) {
	params := arguments.CompanyListArgs{
		ID:        req.Id,
		Name:      req.Name,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}
	companies, err := d.handlers.Company.List(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.CompanyListResponse{}
	for _, company := range companies {
		result.Company = append(result.Company, &pb.Company{
			Id:        company.ID,
			Name:      company.Name,
			Status:    company.Status,
			CreatedBy: company.CreatedBy,
			UpdatedBy: company.UpdatedBy,
		})
	}
	return &result, nil
}

// CompanyGetByID ...
func (d *DeltaServiceImpl) CompanyGetByID(ctx context.Context, req *pb.CompanyGetByIDRequest) (*pb.CompanyGetByIDResponse, error) {
	return nil, nil
}

// CompanyGetByIDs ...
func (d *DeltaServiceImpl) CompanyGetByIDs(ctx context.Context, req *pb.CompanyGetByIDsRequest) (*pb.CompanyGetByIDsResponse, error) {
	return nil, nil
}

// CompanyInsert ...
func (d *DeltaServiceImpl) CompanyInsert(ctx context.Context, req *pb.CompanyInsertRequest) (*pb.CompanyInsertResponse, error) {
	return nil, nil
}

// CompanyUpdate ...
func (d *DeltaServiceImpl) CompanyUpdate(ctx context.Context, req *pb.CompanyUpdateRequest) (*pb.CompanyUpdateResponse, error) {
	return nil, nil
}

// CompanyDelete ...
func (d *DeltaServiceImpl) CompanyDelete(ctx context.Context, req *pb.CompanyDeleteRequest) (*pb.CompanyDeleteResponse, error) {
	return nil, nil
}
