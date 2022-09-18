package domain

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Patronymic        string `json:"patronymic"`
	Role              string `json:"role"`
	EncryptedPassword string `json:"encrypted_password"`
}
