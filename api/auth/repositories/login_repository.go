package repositories

import (
	"notify-backend/api/auth/types"
	"notify-backend/config/database"
	"database/sql"
	

)

func GetUserByUsernameOrEmail(usernameOrEmail string) (types.User, error) {
    var user types.User
    client, err := database.GetDBConnection()
    if err != nil {
        return user, err
    }

    query := `SELECT username, password FROM users WHERE username = $1 OR email = $1`
    row := client.QueryRow(query, usernameOrEmail)

    err = row.Scan(&user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, nil
        }
        return user, err
    }

    return user, nil
}
