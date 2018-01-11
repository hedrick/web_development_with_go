package main

import (
	"fmt"

	"../models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked_dev"
)

// User struct for gorm.Model
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Orders []Order
}

// Order struct for gorm.Model
type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer us.Close()
	// db.LogMode(true)
	us.DestructiveReset()

	// user := models.User{
	// 	Name:     "Michale Scott",
	// 	Email:    "michael@dundermifflin.com",
	// 	Password: "bestboss",
	// }
	// err = us.Create(&user)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", user)
	// if user.Remember == "" {
	// 	panic("Invalid remember token")
	// }

	// user2, err := us.ByRemember(user.Remember)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", *user2)

	// db.AutoMigrate(&User{}, &Order{})

	// name, email := getInfo()
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)

	// 	var user User
	// 	db.Preload("Orders").First(&user)
	// 	if db.Error != nil {
	// 		panic(db.Error)
	// 	}

	// 	fmt.Println("Email:", user.Email)
	// 	fmt.Println("Number of orders:", len(user.Orders))
	// 	fmt.Println("Orders:", user.Orders)
	// }

	// func getInfo() (name, email string) {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Println("What is your name?")
	// 	name, _ = reader.ReadString('\n')
	// 	name = strings.TrimSpace(name)
	// 	fmt.Println("What is your email?")
	// 	email, _ = reader.ReadString('\n')
	// 	email = strings.TrimSpace(email)
	// 	return name, email
	// }

	// func createOrder(db *gorm.DB, user User, amount int, desc string) {
	// 	db.Create(&Order{
	// 		UserID:      user.ID,
	// 		Amount:      amount,
	// 		Description: desc,
	// 	})
	// 	if db.Error != nil {
	// 		panic(db.Error)
	// 	}
}
