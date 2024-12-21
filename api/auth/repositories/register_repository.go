package repositories

import (
	"notify-backend/api/auth/types"
	"notify-backend/api/utils/debug"
	"notify-backend/config/database"
	"database/sql"
)

func RegisterRepository(registerRequest types.RegisterRequest) error {
	client, err := database.GetDBConnection()
	if err != nil {
		debug.LogError(err)
		return err
	}

	// Iniciar una transacción
	tx, err := client.Begin()
	if err != nil {
		debug.LogError(err)
		return err
	}

	// Insertar la información de la compañía
	company := registerRequest.Company
	companyQuery := `INSERT INTO companies 
					(name, phone, address, timezone_id, country_id) 
					VALUES ($1, $2, $3, $4, $5) 
					RETURNING id`
	var companyId int
	err = tx.QueryRow(
		companyQuery, company.Name, company.Phone, 
		company.Address, company.TimezoneId, company.CountryId).Scan(&companyId)
	if err != nil {
		tx.Rollback()
		debug.LogError(err)
		return err
	}

	// Insertar la información del usuario
	user := registerRequest.User
	userQuery := "INSERT INTO users (username, name, email, password, company_id) VALUES ($1, $2, $3, $4, $5)"
	_, err = tx.Exec(userQuery, user.Username, user.Name, user.Email, user.Password, companyId)
	if err != nil {
		tx.Rollback()
		debug.LogError(err)
		return err
	}

	// Confirmar la transacción
	err = tx.Commit()
	if err != nil {
		debug.LogError(err)
		return err
	}

	return nil
}

func ValidateEmailUser(email string) (bool, error) {
	client, err := database.GetDBConnection()
	if err != nil {
		debug.LogError(err)
		return false, err
	}

	query := "SELECT email FROM users WHERE email = $1"
	row := client.QueryRow(query, email)

	var userEmail string
	err = row.Scan(&userEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	return true, nil
}

func ValidateUsernameUser(username string) (bool, error) {
	client, err := database.GetDBConnection()
	if err != nil {
		debug.LogError(err)
		return false, err
	}

	query := "SELECT username FROM users WHERE username = $1"
	row := client.QueryRow(query, username)

	var userUsername string
	err = row.Scan(&userUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	return true, nil
}

func ValidateNameCompany(name string) (bool, error) {
	client, err := database.GetDBConnection()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	query := "SELECT name FROM companies WHERE name = $1"
	row := client.QueryRow(query, name)

	var companyName string
	err = row.Scan(&companyName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	return true, nil
}

func ValidatePhoneCompany(phone string) (bool, error) {
	client, err := database.GetDBConnection()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	query := "SELECT phone FROM companies WHERE phone = $1"
	row := client.QueryRow(query, phone)

	var companyPhone string
	err = row.Scan(&companyPhone)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		debug.LogError(err)
		return false, err
	}

	return true, nil
}
