// @generated
package models

// User ...
type User struct {
	ID        int64  `faker:"unix_time"`
	Email     string `faker:"email"`
	Name      string `faker:"name"`
	CompanyID int64  `faker:"unix_time"`
	Status    string `faker:"name"`
	CreatedBy string `faker:"email"`
	UpdatedBy string `faker:"email"`
}
