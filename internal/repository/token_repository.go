package repository

import (
	"github.com/gocql/gocql"
	"gitlab.com/ft25/iom/engine/am/internal/crypto"
	"gitlab.com/ft25/iom/engine/am/internal/model"
)

type TokenRepository struct {
	session *gocql.Session
	crypto  *crypto.Crypto
}

func NewTokenRepository(session *gocql.Session) *TokenRepository {
	return &TokenRepository{
		session: session,
		crypto:  crypto.New(),
	}
}

func (r *TokenRepository) Create(token *model.Token) (*model.Token, error) {
	// Generate UUID if not provided
	if token.ID == "" {
		token.ID = gocql.TimeUUID().String()
	}

	// Encrypt API key before storing
	encryptedKey, err := r.crypto.Encrypt(token.Secret)
	if err != nil {
		return nil, err
	}
	token.Secret = encryptedKey

	// Begin batch operation
	batch := r.session.NewBatch(gocql.LoggedBatch)

	// Insert into tokens table
	insertTokens := `INSERT INTO tokens (
        groups, id, channel, secret, created_at, updated_at, expires_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	batch.Query(insertTokens,
		token.Groups,
		token.ID,
		token.Channel,
		token.Secret,
		token.CreatedAt,
		token.UpdatedAt,
		token.ExpiresAt,
	)

	// Insert into tokens_by_key table
	insertTokensByName := `INSERT INTO tokens_by_secret (
        groups, channel, id, secret, created_at, updated_at, expires_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	batch.Query(insertTokensByName,
		token.Groups,
		token.Channel,
		token.ID,
		token.Secret,
		token.CreatedAt,
		token.UpdatedAt,
		token.ExpiresAt,
	)

	// Execute batch
	if err := r.session.ExecuteBatch(batch); err != nil {
		return nil, err
	}

	return token, nil
}

func (r *TokenRepository) List() ([]model.Token, error) {
	var tokens []model.Token
	query := `SELECT groups, id, channel, secret, created_at, updated_at, expires_at 
              FROM tokens`

	iter := r.session.Query(query).Iter()
	var token model.Token
	for iter.Scan(
		&token.Groups,
		&token.ID,
		&token.Channel,
		&token.Secret,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.ExpiresAt) {

		// Decrypt API key before returning
		decryptedKey, err := r.crypto.Decrypt(token.Secret)
		if err != nil {
			return nil, err
		}
		token.Secret = decryptedKey
		tokens = append(tokens, token)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *TokenRepository) GetByID(groups string, ID string) (*model.Token, error) {
	var token model.Token
	query := `SELECT groups, id, channel, secret, created_at, updated_at, expires_at 
              FROM tokens WHERE groups = ? and id = ? LIMIT 1`

	if err := r.session.Query(query, groups, ID).Scan(
		&token.Groups,
		&token.ID,
		&token.Channel,
		&token.Secret,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.ExpiresAt); err != nil {
		return nil, err
	}

	// Decrypt API key before returning
	decryptedKey, err := r.crypto.Decrypt(token.Secret)
	if err != nil {
		return nil, err
	}
	token.Secret = decryptedKey

	return &token, nil
}

func (r *TokenRepository) GetByGroups(groups string) ([]model.Token, error) {
	var tokens []model.Token
	query := `SELECT groups, id, channel, secret, created_at, updated_at, expires_at 
			  FROM tokens WHERE groups = ?`

	iter := r.session.Query(query, groups).Iter()
	var token model.Token

	for iter.Scan(
		&token.Groups,
		&token.ID,
		&token.Channel,
		&token.Secret,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.ExpiresAt) {

		// Decrypt API key before returning
		decryptedKey, err := r.crypto.Decrypt(token.Secret)
		if err != nil {
			return nil, err
		}

		token.Secret = decryptedKey
		tokens = append(tokens, token)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *TokenRepository) UpdateSecret(newToken *model.Token, exitToken *model.Token) (*model.Token, error) {

	// Encrypt API key before storing
	encryptedKey, err := r.crypto.Encrypt(newToken.Secret)
	if err != nil {
		return nil, err
	}
	newToken.Secret = encryptedKey

	// Begin batch operation
	batch := r.session.NewBatch(gocql.LoggedBatch)

	// Update tokens table
	updateTokens := `UPDATE tokens SET updated_at = ?, secret = ?
					 WHERE groups = ? and id = ?`

	batch.Query(updateTokens,
		newToken.UpdatedAt,
		newToken.Secret,
		exitToken.Groups,
		exitToken.ID,
	)

	// Delete from tokens_by_secret table
	deleteTokensBySecret := `DELETE FROM tokens_by_secret WHERE groups = ? and secret = ?`
	batch.Query(deleteTokensBySecret, exitToken.Groups, exitToken.Secret)

	// Insert into tokens_by_secret table
	insertTokensBySecret := `INSERT INTO tokens_by_secret (
		groups, channel, id, secret, created_at, updated_at, expires_at
	) VALUES (?, ?, ?, ?, ?, ?, ?)`

	batch.Query(insertTokensBySecret,
		exitToken.Groups,
		exitToken.Channel,
		exitToken.ID,
		exitToken.ExpiresAt,
		newToken.Secret,
		newToken.UpdatedAt,
	)

	// Execute batch
	if err := r.session.ExecuteBatch(batch); err != nil {
		return nil, err
	}

	return newToken, nil
}

func (r *TokenRepository) Delete(groups string, ID string) error {
	// Get existing token to get the name for deletion from tokens_by_key
	_, err := r.GetByID(groups, ID)
	if err != nil {
		return err
	}

	// Begin batch operation
	batch := r.session.NewBatch(gocql.LoggedBatch)

	// Delete from tokens table
	deleteTokens := `DELETE FROM tokens WHERE groups = ? and id = ?`
	batch.Query(deleteTokens, groups, ID)

	// Delete from tokens_by_key table
	deleteTokensBySecret := `DELETE FROM tokens_by_secret WHERE groups = ? and id = ?`
	batch.Query(deleteTokensBySecret, groups, ID)

	// Execute batch
	return r.session.ExecuteBatch(batch)
}

func (r *TokenRepository) GetBySecret(groups string, secret string) (bool, error) {
	// First encrypt the provided API key for comparison
	encryptedSecret, err := r.crypto.Encrypt(secret)
	if err != nil {
		return false, err
	}

	// Check if token exists and is active
	var ID string
	query := `SELECT id FROM tokens_by_secret WHERE groups = ? and secret = ? LIMIT 1`

	if err := r.session.Query(query, groups, encryptedSecret).Scan(&ID); err != nil {
		if err == gocql.ErrNotFound {
			return false, nil
		}
		return false, err
	}

	if ID == "" {
		return false, nil
	}

	return true, nil
}
