// @generated
package models

// Permission ...
type Permission struct {
	ID        int64  `faker:"unix_time"`
	Name      string `faker:"name"`
	Status    string `faker:"name"`
	CreatedBy string `faker:"email"`
	UpdatedBy string `faker:"email"`
}
