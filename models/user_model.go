package models

type User struct {
	ID    uint16   //gorm
	Name  string //gorm
	Email string //gorm
	//Hash  string //gorm
}
