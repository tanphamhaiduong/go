// @generated
package arguments

import "github.com/go-sql-driver/mysql"

// UserGetByID ...
type UserGetByID struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// UserGetByIDs ...
type UserGetByIDs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// UserCount ...
type UserCount struct {
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
	SupervisorID  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleID        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserList ...
type UserList struct {
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
	SupervisorID  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleID        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
	Page          int64          `graphql:"page" validate:"required,min=1"`
	PageSize      int64          `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// UserInsert ...
type UserInsert struct {
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
	SupervisorID  int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleID        int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserUpdate ...
type UserUpdate struct {
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
	SupervisorID  *int64          `graphql:"supervisorId" validate:"omitempty"`
	RoleID        *int64          `graphql:"roleId" validate:"omitempty"`
	CompanyID     *int64          `graphql:"companyId" validate:"omitempty,min=1"`
	Status        *string         `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy     *string         `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy     *string         `graphql:"updatedBy" validate:"omitempty"`
}

// UserDelete ...
type UserDelete struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
