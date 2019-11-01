// @generated
package arguments

// RolePermissionGetByID ...
type RolePermissionGetByID struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// RolePermissionGetByIDs ...
type RolePermissionGetByIDs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// RolePermissionCount ...
type RolePermissionCount struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionList ...
type RolePermissionList struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
	BeginID      int64  `graphql:"beginId" validate:"min=0"`
	EndID        int64  `graphql:"endId" validate:"min=0"`
	PageSize     int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RolePermissionInsert ...
type RolePermissionInsert struct {
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionUpdate ...
type RolePermissionUpdate struct {
	ID           *int64  `graphql:"id" validate:"required,min=1"`
	RoleID       *int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID *int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    *string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionDelete ...
type RolePermissionDelete struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
