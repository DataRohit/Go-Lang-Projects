package main

import (
	"encoding/json"
	"fmt"

	"github.com/datarohit/go-database/pkg/database"
	"github.com/datarohit/go-database/pkg/schemas"
)

const Version = "v1.0.0"

func main() {
	databaseDir := "storage"

	db, err := database.NewDatabase(databaseDir, nil)
	if err != nil {
		fmt.Println("Error initializing the database", err)
	}

	employees := []schemas.User{
		{Name: "John", Age: "23", Contact: "23344333", Company: "Myrl Tech", Address: schemas.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
		{Name: "Paul", Age: "25", Contact: "23344333", Company: "Google", Address: schemas.Address{City: "San Francisco", State: "California", Country: "USA", Pincode: "410013"}},
		{Name: "Robert", Age: "27", Contact: "23344333", Company: "Microsoft", Address: schemas.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
		{Name: "Vince", Age: "29", Contact: "23344333", Company: "Facebook", Address: schemas.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
		{Name: "Neo", Age: "31", Contact: "23344333", Company: "Remote-Teams", Address: schemas.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
		{Name: "Albert", Age: "32", Contact: "23344333", Company: "Dominate", Address: schemas.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, schemas.User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allUsers := []schemas.User{}
	for _, user := range records {
		employeeFound := schemas.User{}
		if err := json.Unmarshal([]byte(user), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println((allUsers))

	// if err := db.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
