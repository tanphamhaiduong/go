// @generated
package arguments

// PermissionGetByIDArgs ...
type PermissionGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// PermissionGetByIDsArgs ...
type PermissionGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// PermissionCountArgs ...
type PermissionCountArgs struct {
	ID        int64  `graphql:"id" validate:"required,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
}

// PermissionListArgs ...
type PermissionListArgs struct {
	ID int64 `graphql:"id" validate:"omitempty,min=1"`

	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// PermissionInsertArgs ...
type PermissionInsertArgs struct {
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
}

// PermissionUpdateArgs ...
type PermissionUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1"`
	Name      *string `graphql:"name" validate:"omitempty"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty`
}

// PermissionDeleteArgs ...
type PermissionDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
