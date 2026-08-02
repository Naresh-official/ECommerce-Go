package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound           = errors.New("User not found")
	ErrEmailAlreadyExists     = errors.New("Email already exists")
	ErrInvalidCurrentPassword = errors.New("Current password is invalid")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProfile(ctx context.Context, userID string) (*ProfileResponse, error) {
	user, err := s.repo.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &ProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  string(user.Role),
	}, nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID string, req UpdateProfileRequest) (*ProfileResponse, error) {
	currentUser, err := s.repo.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if currentUser.Email != req.Email {
		existingUser, err := s.repo.GetUserByEmail(ctx, req.Email)
		if err == nil && existingUser.ID != currentUser.ID {
			return nil, ErrEmailAlreadyExists
		}
		if err != nil && !errors.Is(err, NoRowsError) {
			return nil, err
		}
	}

	updatedUser, err := s.repo.UpdateProfile(ctx, currentUser.ID, req, string(currentUser.Role))
	if err != nil {
		return nil, err
	}

	return &ProfileResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Phone: updatedUser.Phone,
		Role:  string(updatedUser.Role),
	}, nil
}

func (s *Service) ChangePassword(ctx context.Context, userID string, req ChangePasswordRequest) error {
	currentUser, err := s.repo.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return ErrUserNotFound
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.PasswordHash), []byte(req.CurrentPassword)); err != nil {
		return ErrInvalidCurrentPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(ctx, currentUser.ID, string(hashedPassword))
}
