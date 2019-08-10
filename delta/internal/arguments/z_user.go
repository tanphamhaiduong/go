// @generated
package arguments

import "github.com/go-sql-driver/mysql"

// UserGetByIDArgs ...
type UserGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// UserGetByIDsArgs ...
type UserGetByIDsArgs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// UserCountArgs ...
type UserCountArgs struct {
	ID            int64          `graphql:"id" validate:"omitempty,min=1"`
	Username      string         `graphql:"username" validate:"omitempty"`
	Password      string         `graphql:"password" validate:"omitempty"`
	Name          string         `graphql:"name" validate:"omitempty"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty"`
	Reference     string         `graphql:"reference" validate:"omitempty"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty"`
	Extension     string         `graphql:"extension" validate:"omitempty"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty"`
	TelApi        string         `graphql:"telApi" validate:"omitempty"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserListArgs ...
type UserListArgs struct {
	ID            int64          `graphql:"id" validate:"omitempty,min=1"`
	Username      string         `graphql:"username" validate:"omitempty"`
	Password      string         `graphql:"password" validate:"omitempty"`
	Name          string         `graphql:"name" validate:"omitempty"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty"`
	Reference     string         `graphql:"reference" validate:"omitempty"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty"`
	Extension     string         `graphql:"extension" validate:"omitempty"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty"`
	TelApi        string         `graphql:"telApi" validate:"omitempty"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
	Page          int64          `graphql:"page" validate:"required,min=1"`
	PageSize      int64          `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// UserInsertArgs ...
type UserInsertArgs struct {
	Username      string         `graphql:"username" validate:"omitempty"`
	Password      string         `graphql:"password" validate:"omitempty"`
	Name          string         `graphql:"name" validate:"omitempty"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty"`
	Reference     string         `graphql:"reference" validate:"omitempty"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty"`
	Extension     string         `graphql:"extension" validate:"omitempty"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty"`
	TelApi        string         `graphql:"telApi" validate:"omitempty"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserUpdateArgs ...
type UserUpdateArgs struct {
	ID            *int64          `graphql:"id" validate:"required,min=1"`
	Username      *string         `graphql:"username" validate:"omitempty"`
	Password      *string         `graphql:"password" validate:"omitempty"`
	Name          *string         `graphql:"name" validate:"omitempty"`
	DateOfBirth   *mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty"`
	Reference     *string         `graphql:"reference" validate:"omitempty"`
	AvatarUrl     *string         `graphql:"avatarUrl" validate:"omitempty"`
	LicenseNumber *string         `graphql:"licenseNumber" validate:"omitempty"`
	PhoneNumber   *string         `graphql:"phoneNumber" validate:"omitempty"`
	Extension     *string         `graphql:"extension" validate:"omitempty"`
	TelProvider   *string         `graphql:"telProvider" validate:"omitempty"`
	TelApi        *string         `graphql:"telApi" validate:"omitempty"`
	SupervisorId  *int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleId        *int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     *int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        *string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     *string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     *string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserDeleteArgs ...
type UserDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
