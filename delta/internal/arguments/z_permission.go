// @generated
package arguments

// PermissionGetByIDArgs ...
type PermissionGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// PermissionGetByIDsArgs ...
type PermissionGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// PermissionCountArgs ...
type PermissionCountArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// PermissionListArgs ...
type PermissionListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// PermissionInsertArgs ...
type PermissionInsertArgs struct {
	Name      string `graphql:"name" validate:"omitempty" faker:"name"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// PermissionUpdateArgs ...
type PermissionUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1" faker:"unix_time"`
	Name      *string `graphql:"name" validate:"omitempty" faker:"name"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// PermissionDeleteArgs ...
type PermissionDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
