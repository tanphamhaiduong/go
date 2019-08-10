package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission"
	"golang.org/x/crypto/bcrypt"
)

// IRepository ...
type IRepository interface {
	ICoreRepository
	GetByUsername(ctx context.Context, username string) (models.User, error)
}

// HandlerImpl ...
type HandlerImpl struct {
	user           IRepository
	permission     permission.IHandler
	rolePermission rolepermission.IHandler
}

// NewHandler ...
func NewHandler(user IRepository, permission permission.IHandler, rolePermission rolepermission.IHandler) *HandlerImpl {
	return &HandlerImpl{
		user:           user,
		permission:     permission,
		rolePermission: rolePermission,
	}
}

func (h *HandlerImpl) fetchUser(ctx context.Context, param arguments.UserLogin) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
	}).Infof("fetchUser")
	user, err := h.user.GetByUsername(ctx, param.Username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (h *HandlerImpl) fetchRolePermission(ctx context.Context, roleID int64) ([]models.RolePermission, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"roleID":  roleID,
	}).Infof("fetchRolePermission")
	var (
		rolePermissions []models.RolePermission
	)

	rolePermissions, err := h.rolePermission.GetByRoleID(ctx, roleID)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"err":     err,
		}).Errorf("fetchRolePermission cannot fetch rolePermission.GetByRoleID")
		return rolePermissions, err
	}
	return rolePermissions, nil
}

func (h *HandlerImpl) fetchPermission(ctx context.Context, rolePermissions []models.RolePermission) ([]models.Permission, error) {
	logger.WithFields(logger.Fields{
		"TraceID":         ctx.Value("TraceID"),
		"rolePermissions": rolePermissions,
	}).Infof("fetchPermission")
	var (
		permissionIDs []int64
		permissions   []models.Permission
	)
	for _, rolePermission := range rolePermissions {
		permissionIDs = append(permissionIDs, rolePermission.ID)
	}
	permissions, err := h.permission.GetByIDs(ctx, arguments.PermissionGetByIDs{IDs: permissionIDs})
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"err":     err,
		}).Errorf("fetchPermission cannot fetch permission.GetByIDs")
		return permissions, err
	}
	return permissions, nil
}

func (h *HandlerImpl) validateUser(ctx context.Context, user models.User, param arguments.UserLogin) bool {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
		"user":    user,
	}).Infof("validateUser")
	password := []byte(param.Password)
	hashedPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	isValidUser := err == nil
	return isValidUser
}

func (h *HandlerImpl) signTokens(user models.User, permissions []models.Permission, param arguments.UserLogin) (string, error) {
	var (
		secretKey           = os.Getenv("SECRET_KEY")
		jwtKey              = []byte(secretKey)
		expirationTokenTime = time.Now().Add(5 * time.Minute)
		claims              = &models.Claims{
			ID:          user.ID,
			Username:    user.Username,
			CompanyID:   user.CompanyID,
			Permissions: permissions,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTokenTime.Unix(),
			},
		}
		isDevelopmentEnv = secretKey == ""
	)
	if isDevelopmentEnv {
		jwtKey = []byte("test")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// Login ...
func (h *HandlerImpl) Login(ctx context.Context, param arguments.UserLogin) (string, error) {
	logger.WithFields(logger.Fields{
		"TraceID":  ctx.Value("TraceID"),
		"Username": param.Username,
	}).Infof("Login by username and password")
	user, err := h.fetchUser(ctx, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"error":   err,
		}).Errorf("Login cannot fetch user")
		return "", err
	}
	rolePermissions, err := h.fetchRolePermission(ctx, user.RoleID)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"error":   err,
		}).Errorf("Login cannot fetch rolePermissions")
		return "", err
	}
	permissions, err := h.fetchPermission(ctx, rolePermissions)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": ctx.Value("TraceID"),
			"error":   err,
		}).Errorf("Login cannot fetch permissions")
		return "", err
	}
	isValidUser := h.validateUser(ctx, user, param)
	if !isValidUser {
		logger.WithFields(logger.Fields{
			"TraceID":  ctx.Value("TraceID"),
			"Username": param.Username,
		}).Errorf("Login User is not valid")
		return "", errors.New("User is not valid")
	}
	tokenString, err := h.signTokens(user, permissions, param)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID":  ctx.Value("TraceID"),
			"Username": param.Username,
			"Error":    err,
		}).Errorf("Cannot sign token")
		return "", errors.New("Cannot sign token")
	}
	return fmt.Sprintf("Bearer %s", tokenString), nil
}
