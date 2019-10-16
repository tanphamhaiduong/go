// @generated
package arguments

// RoleGetByID ...
type RoleGetByID struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// RoleGetByIDs ...
type RoleGetByIDs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// RoleCount ...
type RoleCount struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleList ...
type RoleList struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
	EndID     int64  `graphql:"beginId" validate:"min=0"`
	BeginID   int64  `graphql:"endId" validate:"min=0"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RoleInsert ...
type RoleInsert struct {
	Name      string `graphql:"name" validate:"omitempty"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleUpdate ...
type RoleUpdate struct {
	ID        *int64  `graphql:"id" validate:"required,min=1"`
	Name      *string `graphql:"name" validate:"omitempty"`
	CompanyID *int64  `graphql:"companyId" validate:"omitempty,min=1"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty"`
}

// RoleDelete ...
type RoleDelete struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
