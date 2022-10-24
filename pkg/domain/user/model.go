package user

type User struct {
	ID        uint64 `json:"id" gorm:"column:id;primaryKey"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email" binding:"email"`
}
