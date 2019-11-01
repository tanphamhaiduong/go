// @generated
package company

import (
	"context"
	"database/sql"
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

func (s *CompanyResolverTestSuite) TestCheckPermission_True() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyGetByID",
				},
			},
		}
		expected = true
	)
	actual := s.Company.checkPermission(*claims, "CompanyGetByID")
	s.Equal(expected, actual)
}

func (s *CompanyResolverTestSuite) TestCheckPermission_False() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyList",
				},
			},
		}
		expected = false
	)
	actual := s.Company.checkPermission(*claims, "CompanyGetByID")
	s.Equal(expected, actual)
}
func (s *CompanyResolverTestSuite) TestGetByID_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByID{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, nil)
	actual, err := s.Company.GetByID(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestGetByID_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByID{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestGetByID_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyGetByID",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByID{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, sql.ErrNoRows)
	actual, err := s.Company.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestGetByID_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		sampleID int64 = 1
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params         = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		company  models.Company
		args     = arguments.CompanyGetByID{
			ID: 1,
		}
	)
	s.MockICompany.On("GetByID", ctx, args).Return(company, sql.ErrNoRows)
	actual, err := s.Company.GetByID(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *CompanyResolverTestSuite) TestList_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyList",
				},
			},
		}
		ctx       = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params    = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"beginId": 1, "endId": 10, "pageSize": 10}}
		companies []models.Company
		args      = arguments.CompanyList{
			BeginID:  1,
			EndID:    10,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, nil)
	actual, err := s.Company.List(params)
	s.Nil(err)
	s.Equal(companies, actual)
}

func (s *CompanyResolverTestSuite) TestList_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyList",
				},
			},
		}
		ctx       = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params    = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"beginId": 1, "endId": 10, "pageSize": 10}}
		companies []models.Company
		args      = arguments.CompanyList{
			BeginID:  1,
			EndID:    10,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestList_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyList",
				},
			},
		}
		ctx       = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params    = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"beginId": 1, "endId": 10, "pageSize": 10}}
		companies []models.Company
		args      = arguments.CompanyList{
			BeginID:  1,
			EndID:    10,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, sql.ErrNoRows)
	actual, err := s.Company.List(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestList_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx       = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params    = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{"beginId": 1, "endId": 10, "pageSize": 10}}
		companies []models.Company
		args      = arguments.CompanyList{
			BeginID:  1,
			EndID:    10,
			PageSize: 10,
		}
	)
	s.MockICompany.On("List", ctx, args).Return(companies, errors.New("some errors"))
	actual, err := s.Company.List(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *CompanyResolverTestSuite) TestCount_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCount{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, nil)
	actual, err := s.Company.Count(params)
	s.Nil(err)
	s.Equal(count, actual)
}

func (s *CompanyResolverTestSuite) TestCount_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCount{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestCount_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyCount",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCount{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, sql.ErrNoRows)
	actual, err := s.Company.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestCount_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx    = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args   = arguments.CompanyCount{}
		count  int64
	)
	s.MockICompany.On("Count", ctx, args).Return(count, errors.New("some errors"))
	actual, err := s.Company.Count(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *CompanyResolverTestSuite) TestInsert_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyInsert",
				},
			},
		}
		ctx     = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params  = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsert{}
		company models.Company
	)
	// Mock Insert
	s.MockICompany.On("Insert", ctx, args).Return(company, nil)
	// Mock GetByID
	s.MockICompany.On("GetByID", ctx, params).Return(company, nil)

	actual, err := s.Company.Insert(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestInsert_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyInsert",
				},
			},
		}
		ctx     = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params  = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsert{}
		company models.Company
	)
	s.MockICompany.On("Insert", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestInsert_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx     = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params  = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsert{}
		company models.Company
	)
	s.MockICompany.On("Insert", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *CompanyResolverTestSuite) TestInsert_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyInsert",
				},
			},
		}
		ctx     = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		params  = graphql.ResolveParams{Context: ctx, Source: map[string]interface{}{}, Args: map[string]interface{}{}}
		args    = arguments.CompanyInsert{}
		company models.Company
	)
	s.MockICompany.On("Insert", ctx, args).Return(company, sql.ErrNoRows)
	actual, err := s.Company.Insert(params)
	s.Nil(actual)
	s.NotNil(err)
}
func (s *CompanyResolverTestSuite) TestUpdate_Success() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdate{
			ID: &sampleID,
		}
		company = models.Company{}
	)
	// Mock for Update
	s.MockICompany.On("Update", ctx, args).Return(company, nil)
	// Mock for LoadByID
	s.MockICompany.On("GetByID", ctx, args).Return(company, nil)
	actual, err := s.Company.Update(params)
	s.Nil(err)
	s.Equal(company, actual)
}

func (s *CompanyResolverTestSuite) TestUpdate_Fail() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdate{
			ID: &sampleID,
		}
		company models.Company
	)
	s.MockICompany.On("Update", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestUpdate_Fail1() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdate{
			ID: &sampleID,
		}
		company models.Company
	)
	s.MockICompany.On("Update", ctx, args).Return(company, errors.New("some errors"))
	actual, err := s.Company.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}

func (s *CompanyResolverTestSuite) TestUpdate_Fail2() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "CompanyUpdate",
				},
			},
		}
		ctx            = context.WithValue(context.Background(), utils.ClaimsKey, claims)
		sampleID int64 = 1
		params         = graphql.ResolveParams{
			Context: ctx,
			Source:  map[string]interface{}{},
			Args:    map[string]interface{}{"id": 1},
		}
		args = arguments.CompanyUpdate{
			ID: &sampleID,
		}
		company models.Company
	)
	s.MockICompany.On("Update", ctx, args).Return(company, sql.ErrNoRows)
	actual, err := s.Company.Update(params)
	s.Nil(actual)
	s.NotNil(err)
}
