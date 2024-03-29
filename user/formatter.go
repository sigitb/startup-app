package user

type UserFormatter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	ImageUrl   string `json:"image_url"`
	Token      string `json:token`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Id:         user.Id,
		Email:      user.Email,
		Name:       user.Name,
		Occupation: user.Occupation,
		ImageUrl:   user.AvatarFileName,
		Token:      token,
	}
	return formatter
}
