package models

type Teacher struct {
	Id        int       `json:"manager_id"`
	UserId    int       `json:"user_id"`
}

type Client struct {
	Id        int       `json:"client_id"`
	UserId    int       `json:"user_id"`
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
	Email    string `json:"email"    binding:"required"`
	Name     string `json:"name"     binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
	Phone    string `json:"phone"`
	Archive  bool   `json:"archive" db:"archive"`
}

type Roles struct {
	Id       int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

type RolesHeaders struct {
	Role string
	Id   int
}