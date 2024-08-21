package repository

import (
	"database/sql"
	"errors"
	"time"

	"nftPlantform/api"  
	"nftPlantform/models"  
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) api.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) CreateUser(username, email, passwordHash, walletAddress string) (int64, error) {
	result, err := r.db.Exec(`
		INSERT INTO users (username, email, password_hash, wallet_address)
		VALUES (?, ?, ?, ?)`,
		username, email, passwordHash, walletAddress)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *MySQLUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, username, email, password_hash, wallet_address, created_at, updated_at
		FROM users WHERE username = ?`, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.WalletAddress, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	var createdAt, updatedAt sql.NullString

	err := r.db.QueryRow(`
		SELECT id, username, email, password_hash, wallet_address, created_at, updated_at
		FROM users WHERE id = ?`, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.WalletAddress, &createdAt, &updatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if createdAt.Valid {
		user.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt.String)
	}
	if updatedAt.Valid {
		user.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt.String)
	}

	return &user, nil
}

func (r *MySQLUserRepository) GetUserByWalletAddress(walletAddress string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, username, email, password_hash, wallet_address, created_at, updated_at
		FROM users WHERE wallet_address = ?`, walletAddress).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.WalletAddress, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) UpdateUser(user *models.User) error {
	_, err := r.db.Exec(`
		UPDATE users SET
			username = ?, email = ?, password_hash = ?, wallet_address = ?, updated_at = ?
		WHERE id = ?`,
		user.Username, user.Email, user.PasswordHash, user.WalletAddress, time.Now(), user.ID)
	return err
}

func (r *MySQLUserRepository) DeleteUser(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}