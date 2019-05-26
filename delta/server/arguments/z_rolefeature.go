// @generated

package arguments

// RoleFeatureGetByIDArgs ...
type RoleFeatureGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// RoleFeatureCountArgs ...
type RoleFeatureCountArgs struct {
	ID        int64  `graphql:"id" validate:"required,min=1"`
	RoleID    int64  `graphql:"roleId" validate:"required,min=1"`
	FeatureID int64  `graphql:"featureId" validate:"required,min=1"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
}

// RoleFeatureListArgs ...
type RoleFeatureListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	RoleID    int64  `graphql:"roleId" validate:"required,min=1"`
	FeatureID int64  `graphql:"featureId" validate:"required,min=1"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// RoleFeatureInsertArgs ...
type RoleFeatureInsertArgs struct {
	RoleID    int64  `graphql:"roleId" validate:"required,min=1"`
	FeatureID int64  `graphql:"featureId" validate:"required,min=1"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
}

// RoleFeatureUpdateArgs ...
type RoleFeatureUpdateArgs struct {
	ID        int64  `graphql:"id" validate:"required,min=1"`
	RoleID    int64  `graphql:"roleId" validate:"required,min=1"`
	FeatureID int64  `graphql:"featureId" validate:"required,min=1"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty`
}

// RoleFeatureDeleteArgs ...
type RoleFeatureDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
