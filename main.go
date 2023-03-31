// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123"
// 	dbname   = "db-go-sql"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

// type Employee struct {
// 	ID        int
// 	Full_name string
// 	Email     string
// 	Age       int
// 	Division  string
// }

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected to database")

// }

// func CreateEmployee() {
// 	var employee = Employee{}

// 	sqlStatement := `
// 	INSERT INTO employees (full_name, email, age, division)
// 	VALUES ($1, $2, $3, $4)
// 	Returning *
// 	`

// 	err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@gmail.com", 23, "IT").
// 		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New Employee Data: %+v \n", employee)
// }

// func GetEmployees() {
// 	var result = []Employee{}

// 	sqlStatement := "SELECT * FROM employees"

// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var employee = Employee{}

// 		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

// 		if err != nil {
// 			panic(err)
// 		}

// 		result = append(result, employee)
// 	}

// 	fmt.Println("Employee datas: ", result)
// }

// func UpdateEmploye() {
// 	sqlStatement := `
// 	UPDATE employees
// 	SET full_name = $2, email = $3, division = $4, age = $5
// 	WHERE id = $1;
// 	`
// 	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airelhidayat@gmail.com", "CurDevs", 24)
// 	if err != nil {
// 		panic(err)
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Updated data amount: ", count)
// }

// func DeleteEmployee() {
// 	sqlStatement := `
// 	DELETE FROM employees
// 	WHERE id = $1;
// 	`
// 	res, err := db.Exec(sqlStatement, 1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Deleted data amount: ", count)
// }

package main

import (
	"challenge-2/database"
	"challenge-2/routers"

	_ "github.com/lib/pq"
)

func main() {
	database.StartDB()

	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}

// func CreateBook(nameBook, author string) {
// 	db := database.GetDB()

// 	Book := models.Book{
// 		NameBook: nameBook,
// 		Author:   author,
// 	}

// 	err := db.Create(&Book).Error

// 	if err != nil {
// 		fmt.Println("Error creating user data: ", err)
// 		return
// 	}

// 	fmt.Println("New Book Data: ", Book)
// }

// func GetBookById(id uint) {
// 	db := database.GetDB()

// 	book := models.Book{}

// 	err := db.First(&book, "id = ?", id).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			fmt.Println("Book data not found")
// 			return
// 		}
// 		print("Error finding book: ", err)
// 	}

// 	fmt.Printf("Book Data: %v \n", book)
// }

// func UpdateBook(id uint, nameBook, author string) {
// 	db := database.GetDB()

// 	book := models.Book{}

// 	err := db.First(&book).Where("id = ?", id).
// 		Updates(models.Book{
// 			NameBook: nameBook,
// 			Author:   author,
// 		}).Error

// 	if err != nil {
// 		fmt.Println("Error updating book data: ", err)
// 		return
// 	}

// 	fmt.Printf("Update book's title: %v \n", book.NameBook)
// }

// func DeleteBook(id uint) {
// 	db := database.GetDB()

// 	book := models.Book{}

// 	err := db.Where("id = ?", id).Delete(&book).Error

// 	if err != nil {
// 		fmt.Println("Error deleting book: ", err.Error())
// 		return
// 	}

// 	fmt.Printf("Product with id %d has been successfully deleted", id)
// }
