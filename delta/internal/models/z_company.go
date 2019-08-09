// @generated
package models

// Company ...
type Company struct {
	ID           int64  `faker:"unix_time"`
	Name         string `faker:"name"`
	CompanyCode  string `faker:"name"`
	Status       string `faker:"word"`
	CreatedBy    string `faker:"email"`
	UpdatedBy    string `faker:"email"`
	ApiSecretKey string `faker:"name"`
}
