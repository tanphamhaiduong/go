// @generated
package arguments

// RolePermissionGetByIDArgs ...
type RolePermissionGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// RolePermissionGetByIDsArgs ...
type RolePermissionGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// RolePermissionCountArgs ...
type RolePermissionCountArgs struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionListArgs ...
type RolePermissionListArgs struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
	Page         int64  `graphql:"page" validate:"required,min=1"`
	PageSize     int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RolePermissionInsertArgs ...
type RolePermissionInsertArgs struct {
	RoleID       int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionUpdateArgs ...
type RolePermissionUpdateArgs struct {
	ID           *int64  `graphql:"id" validate:"required,min=1"`
	RoleID       *int64  `graphql:"roleId" validate:"omitempty,min=1"`
	PermissionID *int64  `graphql:"permissionId" validate:"omitempty,min=1"`
	CreatedBy    *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    *string `graphql:"updatedBy" validate:"omitempty"`
}

// RolePermissionDeleteArgs ...
type RolePermissionDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
