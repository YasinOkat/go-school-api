package repositories

import (
	"database/sql"

	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/utils"
)

func CreateUser(user models.UserCreate) error {
	query := `
	INSERT INTO user
	(username, password, first_name, last_name, phone_number, email, user_type_id, status)
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := utils.DB.Query(
		query,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.UserTypeID,
		user.Status)

	return err
}

func DeleteUser(userID uint) error {
	query := "DELETE FROM user WHERE id = ?"
	_, err := utils.DB.Exec(query, userID)
	return err
}

func GetUsers(filterByActive bool) ([]models.User, error) {
	var query string
	if filterByActive {
		query = "SELECT id, username, password, first_name, last_name, phone_number, email, user_type_id, status FROM user where status = 1"
	} else {
		query = "SELECT id, username, password, first_name, last_name, phone_number, email, user_type_id, status FROM user"
	}
	rows, err := utils.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &user.UserTypeID, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	query := "SELECT id, username, password, first_name, last_name, phone_number, email, user_type_id, status FROM user WHERE id = ?"
	row := utils.DB.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &user.UserTypeID, &user.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, first_name, last_name, phone_number, email, user_type_id, status FROM user WHERE username = ?"
	row := utils.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &user.UserTypeID, &user.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
