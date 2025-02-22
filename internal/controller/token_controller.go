package controller

import (
	"errors"
	"time"

	"gitlab.com/ft25/iom/engine/am/internal/model"
	"gitlab.com/ft25/iom/engine/am/internal/repository"
)

var (
	ErrTokenNotFound    = errors.New("token not found")
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenInactive    = errors.New("token is inactive")
	ErrTokenNameExists  = errors.New("token name already exists")
	ErrInvalidTokenData = errors.New("invalid token data")
)

type TokenService struct {
	tokenRepo *repository.TokenRepository
}

func NewTokenService(tokenRepo *repository.TokenRepository) *TokenService {
	return &TokenService{
		tokenRepo: tokenRepo,
	}
}

func (s *TokenService) Create(token *model.Token) (*model.Token, error) {
	// Validate token data
	if err := s.validateTokenData(token); err != nil {
		return nil, err
	}

	// Set default values
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	if token.ExpiresAt.IsZero() {
		// Default expiration: 1 year
		token.ExpiresAt = time.Now().AddDate(1, 0, 0)
	}

	return s.tokenRepo.Create(token)
}

func (s *TokenService) List() ([]model.Token, error) {
	return s.tokenRepo.List()
}

func (s *TokenService) GetSecrets(groups string) (map[string]string, error) {
	tokens, err := s.GetByGroups(groups)
	if err != nil {
		return nil, err
	}

	secrets := make(map[string]string)
	for _, token := range tokens {
		secrets[token.Secret] = token.Secret
	}

	return secrets, nil
}

func (s *TokenService) GetByGroups(groups string) ([]model.Token, error) {
	tokens, err := s.tokenRepo.GetByGroups(groups)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *TokenService) GetByID(groups string, ID string) (*model.Token, error) {
	token, err := s.tokenRepo.GetByID(groups, ID)
	if err != nil {
		return nil, err
	}

	// Check if token is expired
	if time.Now().After(token.ExpiresAt) {
		return token, ErrTokenExpired
	}

	return token, nil
}

func (s *TokenService) Update(newToken *model.Token, exitToken *model.Token) (*model.Token, error) {

	// Check if token exists
	existing, err := s.tokenRepo.GetByID(exitToken.Groups, exitToken.ID)
	if err != nil {
		return nil, ErrTokenNotFound
	}

	// Update timestamp
	newToken.UpdatedAt = time.Now()

	// Preserve creation time
	newToken.CreatedAt = existing.CreatedAt

	return s.tokenRepo.UpdateSecret(newToken, exitToken)
}

func (s *TokenService) Delete(groups string, ID string) error {
	// Check if token exists
	_, err := s.tokenRepo.GetByID(groups, ID)
	if err != nil {
		return ErrTokenNotFound
	}

	return s.tokenRepo.Delete(groups, ID)
}

func (s *TokenService) GetBySecret(groups string, secret string) error {
	isValid, err := s.tokenRepo.GetBySecret(groups, secret)
	if err != nil {
		return err
	}

	if !isValid {
		return ErrInvalidToken
	}

	return nil
}

/*
func (s *TokenService) ExtendExpiration(id string, duration time.Duration) error {
	token, err := s.tokenRepo.GetByID(id)
	if err != nil {
		return ErrTokenNotFound
	}

	token.ExpiresAt = time.Now().Add(duration)
	token.UpdatedAt = time.Now()

	_, err = s.tokenRepo.Update(token)
	return err
}
*/

func (s *TokenService) validateTokenData(token *model.Token) error {
	if token == nil {
		return ErrInvalidTokenData
	}

	if token.Groups == "" {
		return errors.New("token groups is required")
	}

	if token.Secret == "" {
		return errors.New("token secret is required")
	}

	if !token.ExpiresAt.IsZero() && token.ExpiresAt.Before(time.Now()) {
		return errors.New("expiration date must be in the future")
	}

	return nil
}
