package sqlx

type UserRequest struct {
	ID         uint   `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UserResponse struct {
	ID         uint   `json:"id,omitempty" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name,omitempty" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
}
