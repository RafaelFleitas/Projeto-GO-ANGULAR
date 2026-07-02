package entity

type UserEntity struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Age      int8   `db:"age"`
}
