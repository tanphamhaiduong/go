// @generated
package arguments

// CompanyGetByIDArgs ...
type CompanyGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// CompanyGetByIDsArgs ...
type CompanyGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// CompanyCountArgs ...
type CompanyCountArgs struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name         string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty" faker:"name"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty" faker:"name"`
}

// CompanyListArgs ...
type CompanyListArgs struct {
	ID           int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name         string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty" faker:"name"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty" faker:"name"`
	Page         int64  `graphql:"page" validate:"required,min=1"`
	PageSize     int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// CompanyInsertArgs ...
type CompanyInsertArgs struct {
	Name         string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyCode  string `graphql:"companyCode" validate:"omitempty" faker:"name"`
	Status       string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy    string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy    string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	ApiSecretKey string `graphql:"apiSecretKey" validate:"omitempty" faker:"name"`
}

// CompanyUpdateArgs ...
type CompanyUpdateArgs struct {
	ID           *int64  `graphql:"id" validate:"required,min=1" faker:"unix_time"`
	Name         *string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyCode  *string `graphql:"companyCode" validate:"omitempty" faker:"name"`
	Status       *string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"word"`
	CreatedBy    *string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy    *string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	ApiSecretKey *string `graphql:"apiSecretKey" validate:"omitempty" faker:"name"`
}

// CompanyDeleteArgs ...
type CompanyDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
