package main

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}
