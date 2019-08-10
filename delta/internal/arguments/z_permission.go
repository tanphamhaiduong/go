// @generated
package arguments

// PermissionGetByID ...
type PermissionGetByID struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// PermissionGetByIDs ...
type PermissionGetByIDs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// PermissionCount ...
type PermissionCount struct {
	ID          int64  `graphql:"id" validate:"omitempty,min=1"`
	Name        string `graphql:"name" validate:"omitempty"`
	Description string `graphql:"description" validate:"omitempty"`
	Status      string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy   string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy   string `graphql:"updatedBy" validate:"omitempty"`
}

// PermissionList ...
type PermissionList struct {
	ID          int64  `graphql:"id" validate:"omitempty,min=1"`
	Name        string `graphql:"name" validate:"omitempty"`
	Description string `graphql:"description" validate:"omitempty"`
	Status      string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy   string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy   string `graphql:"updatedBy" validate:"omitempty"`
	Page        int64  `graphql:"page" validate:"required,min=1"`
	PageSize    int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// PermissionInsert ...
type PermissionInsert struct {
	Name        string `graphql:"name" validate:"omitempty"`
	Description string `graphql:"description" validate:"omitempty"`
	Status      string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy   string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy   string `graphql:"updatedBy" validate:"omitempty"`
}

// PermissionUpdate ...
type PermissionUpdate struct {
	ID          *int64  `graphql:"id" validate:"required,min=1"`
	Name        *string `graphql:"name" validate:"omitempty"`
	Description *string `graphql:"description" validate:"omitempty"`
	Status      *string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy   *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy   *string `graphql:"updatedBy" validate:"omitempty"`
}

// PermissionDelete ...
type PermissionDelete struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
