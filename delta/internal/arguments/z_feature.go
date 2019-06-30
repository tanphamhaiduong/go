// @generated
package arguments

// FeatureGetByIDArgs ...
type FeatureGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}

// FeatureGetByIDsArgs ...
type FeatureGetByIDsArgs struct {
	IDs []int64 `graphql:"ids"`
}

// FeatureCountArgs ...
type FeatureCountArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	URL       string `graphql:"url" validate:"omitempty" faker:"name"`
	Meta      string `graphql:"meta" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// FeatureListArgs ...
type FeatureListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1" faker:"unix_time"`
	URL       string `graphql:"url" validate:"omitempty" faker:"name"`
	Meta      string `graphql:"meta" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// FeatureInsertArgs ...
type FeatureInsertArgs struct {
	URL       string `graphql:"url" validate:"omitempty" faker:"name"`
	Meta      string `graphql:"meta" validate:"omitempty" faker:"name"`
	CompanyID int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// FeatureUpdateArgs ...
type FeatureUpdateArgs struct {
	ID        *int64  `graphql:"id" validate:"required,min=1" faker:"unix_time"`
	URL       *string `graphql:"url" validate:"omitempty" faker:"name"`
	Meta      *string `graphql:"meta" validate:"omitempty" faker:"name"`
	CompanyID *int64  `graphql:"companyId" validate:"omitempty,min=1" faker:"unix_time"`
	Status    *string `graphql:"status" validate:"omitempty,oneof=active inactive" faker:"name"`
	CreatedBy *string `graphql:"createdBy" validate:"omitempty" faker:"email"`
	UpdatedBy *string `graphql:"updatedBy" validate:"omitempty" faker:"email"`
}

// FeatureDeleteArgs ...
type FeatureDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1" faker:"unix_time"`
}
