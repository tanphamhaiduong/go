// @generated
package arguments

// RoleGetByIDArgs ...
type RoleGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// RoleGetByIDsArgs ...
type RoleGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// RoleCountArgs ...
type RoleCountArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// RoleListArgs ...
type RoleListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RoleInsertArgs ...
type RoleInsertArgs struct {
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// RoleUpdateArgs ...
type RoleUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1" faker:"unix_time"`
	Name      *string `graphql:"name" validate:"omitempty" faker:"name"`
	CompanyID *int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// RoleDeleteArgs ...
type RoleDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
