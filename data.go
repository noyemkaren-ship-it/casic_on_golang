package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), new(gorm.Config))
	if err != nil {
		panic("Нет базы")
	}
	db.AutoMigrate(&User{})
}

func create_user(name string, password string) {
	var existingUsers []User
	db.Where("name = ?", name).Find(&existingUsers)

	if len(existingUsers) == 0 {
		token, err := GenerateToken(32)
		if err != nil {
			panic("Не смогли создать")
		}
		user := User{Name: name, Password: password, Balance: 0, Token: token}
		db.Create(&user)
	}
}

func login(name string, password string) bool {
	var users []User
	db.Where("name = ? AND password = ?", name, password).Find(&users)

	if len(users) > 0 {
		return true
	}
	return false
}

func change_balance(token string, amount int) string {
	var user User
	result := db.Where("token = ?", token).First(&user)

	if result.Error != nil {
		return "Пользователь не найден"
	}

	user.Balance = user.Balance + amount
	db.Save(&user)

	return fmt.Sprintf("Баланс обновлён. Текущий баланс: %d", user.Balance)
}

func get_balance(token string) int {
	var user User
	result := db.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return 0
	}
	return user.Balance
}

func get_token(name string) string {
	var user User
	result := db.Where("name = ?", name).First(&user)

	if result.Error != nil {
		return "Пользователь не найден"
	}

	return user.Token
}
