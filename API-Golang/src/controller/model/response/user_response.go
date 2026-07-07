package response

//Entrega as informações
type UserResponse struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Age          int8   `json:"age"`
	ProfileImage string `json:"profileImage"`
}

type PaginatedUsersResponse struct {
	Users      []UserResponse `json:"users"`
	Total      int64          `json:"total"`
	Page       int64          `json:"page"`
	PageSize   int64          `json:"pageSize"`
	TotalPages int64          `json:"totalPages"`
}
