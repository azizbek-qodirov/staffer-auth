package models

type RegisterReqSwag struct {
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
}

type RegisterReq struct {
	ID       string `json:"id"`       // User's unique identifier
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
	Role     string
}

type LoginReq struct {
	Email    string `json:"email"`    // User's email
	Password string `json:"password"` // User's password
}

type LoginResp struct {
	Token string `json:"token"` // JWT token provided after successful login
}

type GetProfileReq struct {
	Email string `json:"email"` // Username of the profile to retrieve
}

type GetProfileResp struct {
	ID       string `json:"id"`       // User's unique identifier
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
	Role     string `json:"role"`
}

type GetProfileByIdReq struct {
	ID string `json:"id"` // Username of the profile to retrieve
}

type GetProfileByIdResp struct {
	ID       string `json:"id"`       // User's unique identifier
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Role     string `json:"role"`
}

type ChangeRoleReq struct {
	ID    string `json:"id"`    // Username of the profile to retrieve
	Email string `json:"email"` // Username of the profile to retrieve
	Role  string `json:"role"`  // User's password
}
