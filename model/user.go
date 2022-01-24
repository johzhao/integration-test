package model

type User struct {
	ID   int64  `gorm:"autoIncrement;primaryKey"`
	Name string `gorm:"type:varchar(63);not null"`
	Age  int
}
