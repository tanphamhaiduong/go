// @generated
package models

// Feature ...
type Feature struct {
	ID        int64  `faker:"unix_time"`
	URL       string `faker:"name"`
	Meta      string `faker:"name"`
	CompanyID int64  `faker:"unix_time"`
	Status    string `faker:"name"`
	CreatedBy string `faker:"email"`
	UpdatedBy string `faker:"email"`
}
