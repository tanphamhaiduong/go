// @generated
package models

// Role ...
type Role struct {
	ID        int64  `faker:"unix_time"`
	Name      string `faker:"name"`
	CompanyID int64  `faker:"unix_time"`
	Status    string `faker:"name"`
	CreatedBy string `faker:"email"`
	UpdatedBy string `faker:"email"`
}
