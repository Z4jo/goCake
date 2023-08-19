package entity



type Address struct{
	ID int
	StreetName string 
	HouseNumber int 
	City string 
	PostCode int
}

type User struct {
	ID int 
	Role int8 `gorm:"not null"`
	Username string `gorm:"unique not null"`
	Password string `gorm:"not null"`
	//TODO: hash
	Email string `gorm:"unique not null"` 
	//TODO: hash
	Address Address `gorm:"embedded"`
}
