package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
	Role       string `json:"role"`
	Avatar     string `json:"avatar"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formater := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		Occupation: user.Occupation,
		Role:       user.Role,
		Avatar:     user.Avatar,
		Token:      token,
	}
	return formater
}
