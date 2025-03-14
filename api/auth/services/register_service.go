package services

import (
	"notify-backend/api/auth/repositories"
	"notify-backend/api/auth/types"
	"notify-backend/api/utils/encrypt"
)

func RegisterService(registerRequest types.RegisterRequest) (bool, string, error) {
	// TO DO: hacer la validacion de country y timezone segun el pais por id
	// validate fields company and users
	EmailUserValid, err := repositories.ValidateEmailUser(registerRequest.User.Email)
	if err != nil {
		return false, "", err
	}
	if EmailUserValid {
		return false, "El correo ya existe", nil
	}

	UsernameUserValid, err := repositories.ValidateUsernameUser(registerRequest.User.Username)
	if err != nil {
		return false, "", err
	}
	if UsernameUserValid {
		return false, "El usuario ya existe", nil
	}

	NameComapnyValid, err := repositories.ValidateNameCompany(registerRequest.Company.Name)
	if err != nil {
		return false, "", err
	}
	if NameComapnyValid {
		return false, "Ya existe una compañia con ese nombre", nil
	}

	PhoneCompanyValid, err := repositories.ValidatePhoneCompany(registerRequest.Company.Phone)
	if err != nil {
		return false, "", err
	}
	if PhoneCompanyValid {
		return false, "Una compañia ya tiene asociado ese Telefono", nil
	}

	// Hash password
	hasPassword, err := encrypt.HashPassword(registerRequest.User.Password)
	if err != nil {
		return false, "", err
	}
	registerRequest.User.Password = hasPassword

	errRegi := repositories.RegisterRepository(registerRequest)
	if errRegi != nil {
		return false, "", errRegi
	}

	return true, "User registered successfully", nil
}
