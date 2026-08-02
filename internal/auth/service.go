package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/naresh-official/ecommerce_go/configs"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
	cfg  *configs.JWTConfig
}

var (
	ErrEmailAlreadyExists  = errors.New("Email already exists")
	ErrInvalidCredentials  = errors.New("Invalid Credentials")
	ErrUserNotFound        = errors.New("User not found")
	ErrInvalidRefreshToken = errors.New("Invalid refresh token")
)

func NewService(repo *Repository, cfg *configs.JWTConfig) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest, role string) (*SignUpResponse, error) {
	_, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}

	if !errors.Is(err, NoRowsError) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	user, err := s.repo.CreateUser(ctx, req, role)
	if err != nil {
		return nil, err
	}

	return &SignUpResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *Service) SignIn(ctx context.Context, req SignInRequest, role string) (*SignInServiceResult, error) {
	user, err := s.repo.GetUserByEmailAndRole(ctx, req.Email, role)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := GenerateAccessToken(s.cfg, user.ID.String(), user.Email, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken(s.cfg, user.ID.String(), role)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateRefreshToken(ctx, user.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &SignInServiceResult{
		User: SignInResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) GetMeUser(ctx context.Context, userID string) (*GetMeResponse, error) {
	user, err := s.repo.GetUserById(ctx, uuid.MustParse(userID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &GetMeResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil
}

func (s *Service) UpdateAccessToken(ctx context.Context, refreshToken string) (*UpdateAccessTokenServiceResult, error) {
	refreshTokenClaims, err := ValidateRefreshToken(s.cfg, refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserById(ctx, uuid.MustParse(refreshTokenClaims.UserID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if !user.RefreshToken.Valid || user.RefreshToken.String != refreshToken {
		return nil, ErrInvalidRefreshToken
	}

	newAccessToken, err := GenerateAccessToken(s.cfg, user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := GenerateRefreshToken(s.cfg, user.ID.String(), string(user.Role))
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateRefreshToken(ctx, user.ID, newRefreshToken)
	if err != nil {
		return nil, err
	}
	return &UpdateAccessTokenServiceResult{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
