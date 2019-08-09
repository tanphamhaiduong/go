// @generated
package models

import "github.com/go-sql-driver/mysql"

// User ...
type User struct {
	ID            int64
	Username      string
	Password      string
	Name          string
	DateOfBirth   mysql.NullTime
	Reference     string
	AvatarUrl     string
	LicenseNumber string
	PhoneNumber   string
	Extension     string
	TelProvider   string
	TelApi        string
	SupervisorId  int64
	RoleId        int64
	CompanyID     int64
	Status        string
	CreatedBy     string
	UpdatedBy     string
}
