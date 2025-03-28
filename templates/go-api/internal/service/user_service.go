package service

import (
	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(req *models.CreateUserRequest) (*models.UserResponse, error) {
	// Vérifier si l'email existe déjà
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return s.toResponse(user), nil
}

func (s *UserService) GetByID(id uint) (*models.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(user), nil
}

func (s *UserService) GetAll() ([]models.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, *s.toResponse(&user))
	}
	return responses, nil
}

func (s *UserService) Update(id uint, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return s.toResponse(user), nil
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) toResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
} 