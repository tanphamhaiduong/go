// @generated
package arguments

// UserGetByIDArgs ...
type UserGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// UserGetByIDsArgs ...
type UserGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// UserCountArgs ...
type UserCountArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Email     string `graphql:"email" validate:"omitempty" `
	Name      string `graphql:"name" validate:"omitempty" `
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" `
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" `
	CreatedBy string `graphql:"createdBy" validate:"omitempty `
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty `
}

// UserListArgs ...
type UserListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Email     string `graphql:"email" validate:"omitempty" `
	Name      string `graphql:"name" validate:"omitempty" `
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" `
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" `
	CreatedBy string `graphql:"createdBy" validate:"omitempty `
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty `
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// UserInsertArgs ...
type UserInsertArgs struct {
	Email     string `graphql:"email" validate:"omitempty" `
	Name      string `graphql:"name" validate:"omitempty" `
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" `
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" `
	CreatedBy string `graphql:"createdBy" validate:"omitempty `
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty `
}

// UserUpdateArgs ...
type UserUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1" `
	Email     *string `graphql:"email" validate:"omitempty" `
	Name      *string `graphql:"name" validate:"omitempty" `
	CompanyID *int64  `graphql:"companyId" validate:"omitempty,min=1" `
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive" `
	CreatedBy *string `graphql:"createdBy" validate:"omitempty `
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty `
}

// UserDeleteArgs ...
type UserDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
