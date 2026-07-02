package response

//Entrega as informações
type UserResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
