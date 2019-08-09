// @generated
package models

import "github.com/go-sql-driver/mysql"

// User ...
type User struct {
	ID            int64          `faker:"unix_time"`
	Username      string         `faker:"username"`
	Password      string         `faker:"password"`
	Name          string         `faker:"name"`
	DateOfBirth   mysql.NullTime `faker:"date"`
	Reference     string         `faker:"name"`
	AvatarUrl     string         `faker:"domain_name"`
	LicenseNumber string         `faker:"phone_number"`
	PhoneNumber   string         `faker:"phone_number"`
	Extension     string         `faker:"phone_number"`
	TelProvider   string         `faker:"phone_number"`
	TelApi        string         `faker:"domain_name"`
	SupervisorId  int64          `faker:"unix_time"`
	RoleId        int64          `faker:"unix_time"`
	CompanyID     int64          `faker:"unix_time"`
	Status        string         `faker:"name"`
	CreatedBy     string         `faker:"email"`
	UpdatedBy     string         `faker:"email"`
}
