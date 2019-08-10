// @generated
package arguments

// RoleGetByIDArgs ...
type RoleGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// RoleGetByIDsArgs ...
type RoleGetByIDsArgs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// RoleCountArgs ...
type RoleCountArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleListArgs ...
type RoleListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RoleInsertArgs ...
type RoleInsertArgs struct {
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleUpdateArgs ...
type RoleUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1"`
	Name      *string `graphql:"name" validate:"omitempty"`
	CompanyID *int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleDeleteArgs ...
type RoleDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
