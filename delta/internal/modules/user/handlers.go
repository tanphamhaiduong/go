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
	"golang.org/x/crypto/bcrypt"
)

// IRepository ...
type IRepository interface {
	ICoreRepository
	GetByUsername(ctx context.Context, param arguments.UserGetByUsername) (models.User, error)
}

// HandlerImpl ...
type HandlerImpl struct {
	user IRepository
}

// NewHandler ...
func NewHandler(user IRepository) *HandlerImpl {
	return &HandlerImpl{
		user: user,
	}
}

func (h *HandlerImpl) fetchUser(ctx context.Context, param arguments.UserLogin) (models.User, error) {
	logger.WithFields(logger.Fields{
		"TraceID": ctx.Value("TraceID"),
		"param":   param,
	}).Infof("fetchUser")
	getByUsernameArgs := arguments.UserGetByUsername{Username: param.Username}
	user, err := h.user.GetByUsername(ctx, getByUsernameArgs)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
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

func (h *HandlerImpl) signTokens(user models.User, param arguments.UserLogin) (string, error) {
	var (
		secretKey           = os.Getenv("SECRET_KEY")
		jwtKey              = []byte(secretKey)
		expirationTokenTime = time.Now().Add(5 * time.Minute)
		claims              = &models.Claims{
			User: user,
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
		return "", err
	}
	isValidUser := h.validateUser(ctx, user, param)
	if !isValidUser {
		logger.WithFields(logger.Fields{
			"TraceID":  ctx.Value("TraceID"),
			"Username": param.Username,
		}).Errorf("User is not valid")
		return "", errors.New("User is not valid")
	}
	tokenString, err := h.signTokens(user, param)
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
