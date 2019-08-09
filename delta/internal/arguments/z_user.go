// @generated
package arguments

import "github.com/go-sql-driver/mysql"

// UserGetByIDArgs ...
type UserGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// UserGetByIDsArgs ...
type UserGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// UserCountArgs ...
type UserCountArgs struct {
	ID            int64          `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Username      string         `graphql:"username" validate:"omitempty" faker:"username"`
	Password      string         `graphql:"password" validate:"omitempty" faker:"password"`
	Name          string         `graphql:"name" validate:"omitempty" faker:"name"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty" `
	Reference     string         `graphql:"reference" validate:"omitempty" faker:"name"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty" faker:"domain_name"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty" faker:"phone_number"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty" faker:"phone_number"`
	Extension     string         `graphql:"extension" validate:"omitempty" faker:"phone_number"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty" faker:"phone_number"`
	TelApi        string         `graphql:"telApi" validate:"omitempty" faker:"domain_name"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty" faker:"unix_time"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty" faker:"unix_time"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// UserListArgs ...
type UserListArgs struct {
	ID            int64          `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Username      string         `graphql:"username" validate:"omitempty" faker:"username"`
	Password      string         `graphql:"password" validate:"omitempty" faker:"password"`
	Name          string         `graphql:"name" validate:"omitempty" faker:"name"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty" `
	Reference     string         `graphql:"reference" validate:"omitempty" faker:"name"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty" faker:"domain_name"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty" faker:"phone_number"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty" faker:"phone_number"`
	Extension     string         `graphql:"extension" validate:"omitempty" faker:"phone_number"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty" faker:"phone_number"`
	TelApi        string         `graphql:"telApi" validate:"omitempty" faker:"domain_name"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty" faker:"unix_time"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty" faker:"unix_time"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	Page          int64          `graphql:"page" validate:"required,min=1"`
	PageSize      int64          `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// UserInsertArgs ...
type UserInsertArgs struct {
	Username      string         `graphql:"username" validate:"omitempty" faker:"username"`
	Password      string         `graphql:"password" validate:"omitempty" faker:"password"`
	Name          string         `graphql:"name" validate:"omitempty" faker:"name"`
	DateOfBirth   mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty" `
	Reference     string         `graphql:"reference" validate:"omitempty" faker:"name"`
	AvatarUrl     string         `graphql:"avatarUrl" validate:"omitempty" faker:"domain_name"`
	LicenseNumber string         `graphql:"licenseNumber" validate:"omitempty" faker:"phone_number"`
	PhoneNumber   string         `graphql:"phoneNumber" validate:"omitempty" faker:"phone_number"`
	Extension     string         `graphql:"extension" validate:"omitempty" faker:"phone_number"`
	TelProvider   string         `graphql:"telProvider" validate:"omitempty" faker:"phone_number"`
	TelApi        string         `graphql:"telApi" validate:"omitempty" faker:"domain_name"`
	SupervisorId  int64          `graphql:"supervisorId" validate:"omitempty" faker:"unix_time"`
	RoleId        int64          `graphql:"roleId" validate:"omitempty" faker:"unix_time"`
	CompanyID     int64          `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status        string         `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy     string         `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy     string         `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// UserUpdateArgs ...
type UserUpdateArgs struct {
	ID            *int64          `graphql:"id" validate:"required,min=1" faker:"unix_time"`
	Username      *string         `graphql:"username" validate:"omitempty" faker:"username"`
	Password      *string         `graphql:"password" validate:"omitempty" faker:"password"`
	Name          *string         `graphql:"name" validate:"omitempty" faker:"name"`
	DateOfBirth   *mysql.NullTime `graphql:"dateOfBirth" validate:"omitempty" `
	Reference     *string         `graphql:"reference" validate:"omitempty" faker:"name"`
	AvatarUrl     *string         `graphql:"avatarUrl" validate:"omitempty" faker:"domain_name"`
	LicenseNumber *string         `graphql:"licenseNumber" validate:"omitempty" faker:"phone_number"`
	PhoneNumber   *string         `graphql:"phoneNumber" validate:"omitempty" faker:"phone_number"`
	Extension     *string         `graphql:"extension" validate:"omitempty" faker:"phone_number"`
	TelProvider   *string         `graphql:"telProvider" validate:"omitempty" faker:"phone_number"`
	TelApi        *string         `graphql:"telApi" validate:"omitempty" faker:"domain_name"`
	SupervisorId  *int64          `graphql:"supervisorId" validate:"omitempty" faker:"unix_time"`
	RoleId        *int64          `graphql:"roleId" validate:"omitempty" faker:"unix_time"`
	CompanyID     *int64          `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status        *string         `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy     *string         `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy     *string         `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// UserDeleteArgs ...
type UserDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
