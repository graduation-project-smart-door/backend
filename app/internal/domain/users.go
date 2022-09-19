package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                string `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Patronymic        string `json:"patronymic"`
	Role              string `json:"role"`
	EncryptedPassword string `json:"encrypted_password"`
}

type CreateUser struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Role       string `json:"role"`
	Password   string `json:"password"`
}

func (u *User) FromCreateUser(model CreateUser) {
	u.Name = model.Name
	u.Surname = model.Surname
	u.Patronymic = model.Patronymic
	u.Email = model.Email
	u.Role = model.Role
}

func EncryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
