package schema

type SignUpSchemaIn struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type SignInSchemaIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserMeSchema struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsSuperuser bool   `json:"is_superuser"`
}
