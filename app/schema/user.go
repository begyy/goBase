package schema

type SignUpSchemaIn struct {
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
}

type SignInSchemaIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserMeSchema struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsSuperuser bool   `json:"is_superuser"`
}

type UserSchema struct {
	UserMeSchema
	Password string `json:"password"`
}
