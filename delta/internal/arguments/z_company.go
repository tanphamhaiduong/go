// @generated
package arguments

// CompanyGetByID ...
type CompanyGetByID struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// CompanyGetByIDs ...
type CompanyGetByIDs struct {
	IDs []int64 `graphql:"ids" validate:"required"`
}

// CompanyCount ...
type CompanyCount struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	Name         string `graphql:"name" validate:"omitempty"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty"`
}

// CompanyList ...
type CompanyList struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1"`
	Name         string `graphql:"name" validate:"omitempty"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty"`
	Page         int64  `graphql:"page" validate:"required,min=1"`
	PageSize     int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// CompanyInsert ...
type CompanyInsert struct {
	Name         string `graphql:"name" validate:"omitempty"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty"`
}

// CompanyUpdate ...
type CompanyUpdate struct {
	ID           *int64  `graphql:"id" validate:"required,min=1"`
	Name         *string `graphql:"name" validate:"omitempty"`
	CompanyCode  *string `graphql:"companyCode" validate:"omitempty"`
	Status       *string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy    *string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy    *string `graphql:"updatedBy" validate:"omitempty"`
	ApiSecretKey *string `graphql:"apiSecretKey" validate:"omitempty"`
}

// CompanyDelete ...
type CompanyDelete struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
