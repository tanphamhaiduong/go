// @generated
package models

// RolePermission ...
type RolePermission struct {
	ID           int64  `faker:"unix_time"`
	RoleID       int64  `faker:"unix_time"`
	PermissionID int64  `faker:"unix_time"`
	CreatedBy    string `faker:"email"`
	UpdatedBy    string `faker:"email"`
}
