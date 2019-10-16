// @generated
package grpcservice

import (
	"context"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/pb"
)

// CompanyList ...
func (d *DeltaServiceImpl) CompanyList(ctx context.Context, req *pb.CompanyListRequest) (*pb.CompanyListResponse, error) {
	params := arguments.CompanyList{
		ID:        req.Id,
		Name:      req.Name,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
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
	params := arguments.CompanyGetByID{
		ID: req.Id,
	}
	company, err := d.handlers.Company.GetByID(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.CompanyGetByIDResponse{
		Company: &pb.Company{
			Id:        company.ID,
			Name:      company.Name,
			Status:    company.Status,
			CreatedBy: company.CreatedBy,
			UpdatedBy: company.UpdatedBy,
		},
	}
	return &result, nil
}

// CompanyInsert ...
func (d *DeltaServiceImpl) CompanyInsert(ctx context.Context, req *pb.CompanyInsertRequest) (*pb.CompanyInsertResponse, error) {
	params := arguments.CompanyInsert{
		Name:      req.Name,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
	}
	company, err := d.handlers.Company.Insert(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.CompanyInsertResponse{
		Company: &pb.Company{
			Id:        company.ID,
			Name:      company.Name,
			Status:    company.Status,
			CreatedBy: company.CreatedBy,
			UpdatedBy: company.UpdatedBy,
		},
	}
	return &result, nil
}

// CompanyUpdate ...
func (d *DeltaServiceImpl) CompanyUpdate(ctx context.Context, req *pb.CompanyUpdateRequest) (*pb.CompanyUpdateResponse, error) {
	params := arguments.CompanyUpdate{
		ID:        &req.Id,
		Name:      &req.Name,
		Status:    &req.Status,
		CreatedBy: &req.CreatedBy,
		UpdatedBy: &req.UpdatedBy,
	}
	company, err := d.handlers.Company.Update(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.CompanyUpdateResponse{
		Company: &pb.Company{
			Id:        company.ID,
			Name:      company.Name,
			Status:    company.Status,
			CreatedBy: company.CreatedBy,
			UpdatedBy: company.UpdatedBy,
		},
	}
	return &result, nil
}

// PermissionList ...
func (d *DeltaServiceImpl) PermissionList(ctx context.Context, req *pb.PermissionListRequest) (*pb.PermissionListResponse, error) {
	params := arguments.PermissionList{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
		PageSize:    req.PageSize,
	}
	permissions, err := d.handlers.Permission.List(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.PermissionListResponse{}
	for _, permission := range permissions {
		result.Permission = append(result.Permission, &pb.Permission{
			Id:          permission.ID,
			Name:        permission.Name,
			Description: permission.Description,
			Status:      permission.Status,
			CreatedBy:   permission.CreatedBy,
			UpdatedBy:   permission.UpdatedBy,
		})
	}
	return &result, nil
}

// PermissionGetByID ...
func (d *DeltaServiceImpl) PermissionGetByID(ctx context.Context, req *pb.PermissionGetByIDRequest) (*pb.PermissionGetByIDResponse, error) {
	params := arguments.PermissionGetByID{
		ID: req.Id,
	}
	permission, err := d.handlers.Permission.GetByID(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.PermissionGetByIDResponse{
		Permission: &pb.Permission{
			Id:          permission.ID,
			Name:        permission.Name,
			Description: permission.Description,
			Status:      permission.Status,
			CreatedBy:   permission.CreatedBy,
			UpdatedBy:   permission.UpdatedBy,
		},
	}
	return &result, nil
}

// PermissionInsert ...
func (d *DeltaServiceImpl) PermissionInsert(ctx context.Context, req *pb.PermissionInsertRequest) (*pb.PermissionInsertResponse, error) {
	params := arguments.PermissionInsert{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}
	permission, err := d.handlers.Permission.Insert(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.PermissionInsertResponse{
		Permission: &pb.Permission{
			Id:          permission.ID,
			Name:        permission.Name,
			Description: permission.Description,
			Status:      permission.Status,
			CreatedBy:   permission.CreatedBy,
			UpdatedBy:   permission.UpdatedBy,
		},
	}
	return &result, nil
}

// RoleList ...
func (d *DeltaServiceImpl) RoleList(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListResponse, error) {
	params := arguments.RoleList{
		ID:        req.Id,
		Name:      req.Name,
		CompanyID: req.CompanyId,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
		PageSize:  req.PageSize,
	}
	roles, err := d.handlers.Role.List(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.RoleListResponse{}
	for _, role := range roles {
		result.Role = append(result.Role, &pb.Role{
			Id:        role.ID,
			Name:      role.Name,
			CompanyId: role.CompanyID,
			Status:    role.Status,
			CreatedBy: role.CreatedBy,
			UpdatedBy: role.UpdatedBy,
		})
	}
	return &result, nil
}

// RoleGetByID ...
func (d *DeltaServiceImpl) RoleGetByID(ctx context.Context, req *pb.RoleGetByIDRequest) (*pb.RoleGetByIDResponse, error) {
	params := arguments.RoleGetByID{
		ID: req.Id,
	}
	role, err := d.handlers.Role.GetByID(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.RoleGetByIDResponse{
		Role: &pb.Role{
			Id:        role.ID,
			Name:      role.Name,
			CompanyId: role.CompanyID,
			Status:    role.Status,
			CreatedBy: role.CreatedBy,
			UpdatedBy: role.UpdatedBy,
		},
	}
	return &result, nil
}

// RoleInsert ...
func (d *DeltaServiceImpl) RoleInsert(ctx context.Context, req *pb.RoleInsertRequest) (*pb.RoleInsertResponse, error) {
	params := arguments.RoleInsert{
		Name:      req.Name,
		CompanyID: req.CompanyId,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
	}
	role, err := d.handlers.Role.Insert(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.RoleInsertResponse{
		Role: &pb.Role{
			Id:        role.ID,
			Name:      role.Name,
			CompanyId: role.CompanyID,
			Status:    role.Status,
			CreatedBy: role.CreatedBy,
			UpdatedBy: role.UpdatedBy,
		},
	}
	return &result, nil
}

// RoleUpdate ...
func (d *DeltaServiceImpl) RoleUpdate(ctx context.Context, req *pb.RoleUpdateRequest) (*pb.RoleUpdateResponse, error) {
	params := arguments.RoleUpdate{
		ID:        &req.Id,
		Name:      &req.Name,
		CompanyID: &req.CompanyId,
		Status:    &req.Status,
		CreatedBy: &req.CreatedBy,
		UpdatedBy: &req.UpdatedBy,
	}
	role, err := d.handlers.Role.Update(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.RoleUpdateResponse{
		Role: &pb.Role{
			Id:        role.ID,
			Name:      role.Name,
			CompanyId: role.CompanyID,
			Status:    role.Status,
			CreatedBy: role.CreatedBy,
			UpdatedBy: role.UpdatedBy,
		},
	}
	return &result, nil
}

// UserList ...
func (d *DeltaServiceImpl) UserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	dateOfBirth, err := time.Parse(time.RFC3339, req.DateOfBirth)
	if err != nil {
		return nil, err
	}
	params := arguments.UserList{
		ID:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		DateOfBirth: mysql.NullTime{
			Time:  dateOfBirth,
			Valid: true,
		},
		Reference:     req.Reference,
		AvatarUrl:     req.AvatarUrl,
		LicenseNumber: req.LicenseNumber,
		PhoneNumber:   req.PhoneNumber,
		Extension:     req.Extension,
		TelProvider:   req.TelProvider,
		TelApi:        req.TelApi,
		SupervisorID:  req.SupervisorId,
		RoleID:        req.RoleId,
		CompanyID:     req.CompanyId,
		Status:        req.Status,
		CreatedBy:     req.CreatedBy,
		UpdatedBy:     req.UpdatedBy,
		PageSize:      req.PageSize,
	}
	users, err := d.handlers.User.List(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.UserListResponse{}
	for _, user := range users {
		result.User = append(result.User, &pb.User{
			Id:            user.ID,
			Username:      user.Username,
			Password:      user.Password,
			Name:          user.Name,
			DateOfBirth:   user.DateOfBirth.Time.String(),
			Reference:     user.Reference,
			AvatarUrl:     user.AvatarUrl,
			LicenseNumber: user.LicenseNumber,
			PhoneNumber:   user.PhoneNumber,
			Extension:     user.Extension,
			TelProvider:   user.TelProvider,
			TelApi:        user.TelApi,
			SupervisorId:  user.SupervisorID,
			RoleId:        user.RoleID,
			CompanyId:     user.CompanyID,
			Status:        user.Status,
			CreatedBy:     user.CreatedBy,
			UpdatedBy:     user.UpdatedBy,
		})
	}
	return &result, nil
}

// UserGetByID ...
func (d *DeltaServiceImpl) UserGetByID(ctx context.Context, req *pb.UserGetByIDRequest) (*pb.UserGetByIDResponse, error) {
	params := arguments.UserGetByID{
		ID: req.Id,
	}
	user, err := d.handlers.User.GetByID(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.UserGetByIDResponse{
		User: &pb.User{
			Id:            user.ID,
			Username:      user.Username,
			Password:      user.Password,
			Name:          user.Name,
			DateOfBirth:   user.DateOfBirth.Time.String(),
			Reference:     user.Reference,
			AvatarUrl:     user.AvatarUrl,
			LicenseNumber: user.LicenseNumber,
			PhoneNumber:   user.PhoneNumber,
			Extension:     user.Extension,
			TelProvider:   user.TelProvider,
			TelApi:        user.TelApi,
			SupervisorId:  user.SupervisorID,
			RoleId:        user.RoleID,
			CompanyId:     user.CompanyID,
			Status:        user.Status,
			CreatedBy:     user.CreatedBy,
			UpdatedBy:     user.UpdatedBy,
		},
	}
	return &result, nil
}

// UserInsert ...
func (d *DeltaServiceImpl) UserInsert(ctx context.Context, req *pb.UserInsertRequest) (*pb.UserInsertResponse, error) {
	dateOfBirth, err := time.Parse(time.RFC3339, req.DateOfBirth)
	if err != nil {
		return nil, err
	}
	params := arguments.UserInsert{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		DateOfBirth: mysql.NullTime{
			Time:  dateOfBirth,
			Valid: true,
		},
		Reference:     req.Reference,
		AvatarUrl:     req.AvatarUrl,
		LicenseNumber: req.LicenseNumber,
		PhoneNumber:   req.PhoneNumber,
		Extension:     req.Extension,
		TelProvider:   req.TelProvider,
		TelApi:        req.TelApi,
		SupervisorID:  req.SupervisorId,
		RoleID:        req.RoleId,
		CompanyID:     req.CompanyId,
		Status:        req.Status,
		CreatedBy:     req.CreatedBy,
		UpdatedBy:     req.UpdatedBy,
	}
	user, err := d.handlers.User.Insert(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.UserInsertResponse{
		User: &pb.User{
			Id:            user.ID,
			Username:      user.Username,
			Password:      user.Password,
			Name:          user.Name,
			DateOfBirth:   user.DateOfBirth.Time.String(),
			Reference:     user.Reference,
			AvatarUrl:     user.AvatarUrl,
			LicenseNumber: user.LicenseNumber,
			PhoneNumber:   user.PhoneNumber,
			Extension:     user.Extension,
			TelProvider:   user.TelProvider,
			TelApi:        user.TelApi,
			SupervisorId:  user.SupervisorID,
			RoleId:        user.RoleID,
			CompanyId:     user.CompanyID,
			Status:        user.Status,
			CreatedBy:     user.CreatedBy,
			UpdatedBy:     user.UpdatedBy,
		},
	}
	return &result, nil
}

// UserUpdate ...
func (d *DeltaServiceImpl) UserUpdate(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	dateOfBirth, err := time.Parse(time.RFC3339, req.DateOfBirth)
	if err != nil {
		return nil, err
	}
	params := arguments.UserUpdate{
		ID:       &req.Id,
		Username: &req.Username,
		Password: &req.Password,
		Name:     &req.Name,
		DateOfBirth: &mysql.NullTime{
			Time:  dateOfBirth,
			Valid: true,
		},
		Reference:     &req.Reference,
		AvatarUrl:     &req.AvatarUrl,
		LicenseNumber: &req.LicenseNumber,
		PhoneNumber:   &req.PhoneNumber,
		Extension:     &req.Extension,
		TelProvider:   &req.TelProvider,
		TelApi:        &req.TelApi,
		SupervisorID:  &req.SupervisorId,
		RoleID:        &req.RoleId,
		CompanyID:     &req.CompanyId,
		Status:        &req.Status,
		CreatedBy:     &req.CreatedBy,
		UpdatedBy:     &req.UpdatedBy,
	}
	user, err := d.handlers.User.Update(ctx, params)
	if err != nil {
		return nil, err
	}
	result := pb.UserUpdateResponse{
		User: &pb.User{
			Id:            user.ID,
			Username:      user.Username,
			Password:      user.Password,
			Name:          user.Name,
			DateOfBirth:   user.DateOfBirth.Time.String(),
			Reference:     user.Reference,
			AvatarUrl:     user.AvatarUrl,
			LicenseNumber: user.LicenseNumber,
			PhoneNumber:   user.PhoneNumber,
			Extension:     user.Extension,
			TelProvider:   user.TelProvider,
			TelApi:        user.TelApi,
			SupervisorId:  user.SupervisorID,
			RoleId:        user.RoleID,
			CompanyId:     user.CompanyID,
			Status:        user.Status,
			CreatedBy:     user.CreatedBy,
			UpdatedBy:     user.UpdatedBy,
		},
	}
	return &result, nil
}
