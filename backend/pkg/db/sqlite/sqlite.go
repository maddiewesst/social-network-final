package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	// "runtime"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite" // Correct import
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file" // File-based migrations
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	NickName  string `json:"nickName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
	Image     string `json:"image"`
	About     string `json:"about"`
	Public    bool   `json:"public"`
}

type Post struct {
	ID        int    `json:"id"`
	Author    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	Image     string `json:"image"`
	Privacy   bool   `json:"privacy"`
}

// var db *sql.DB

// var m *migrate.Migrate

// run migrations

func RunMigration() *migrate.Migrate {

	cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Current working directory:", cwd)


	migrationPath := "file://" + filepath.Join("pkg", "db", "migration", "sqlite")
	dbPath := filepath.Join("pkg", "db", "database.db")

	if runtime.GOOS == "darwin" {
		migrationPath = "file://" + filepath.Join("..", "pkg", "db", "migration", "sqlite")
		dbPath = filepath.Join("..", "pkg", "db", "database.db")
	}

	// Check if the database file exists
	// if _, err := os.Stat(dbPath); os.IsNotExist(err) {
	// 	// If it doesn't exist, log that the file will be created
	// 	log.Println("Database file not found, creating it...")
	// }

	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create the database driver instance
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}

	// Initialize migration
	m, err := migrate.NewWithDatabaseInstance(migrationPath, "sqlite3", driver)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration applied successfully.")
	return m
}

// RemoveMigration rolls back to a specific migration version (e.g., 0 to undo all migrations)
func RemoveMigration(m *migrate.Migrate, version uint) {
	if err := m.Migrate(version); err != nil {
		log.Fatalf("Error rolling back migrations: %v", err)
	}
	fmt.Printf("Rolled back to version %d\n", version)
}

func DbConnect() *sql.DB {

	dbPath := filepath.Join("pkg", "db", "database.db")

	if runtime.GOOS == "darwin" {
		dbPath = filepath.Join("..", "pkg", "db", "database.db")
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	return db
}

// insert mock user data

func InsertMockUserData() {

	// fetch api for mock data

	var res *http.Response

	res, _ = http.Get("https://63f35a0e864fb1d60014de90.mockapi.io/users")

	resData, _ := io.ReadAll(res.Body)

	// Unmarshall http response

	var responseObject []User

	json.Unmarshal(resData, &responseObject)

	// insert into database

	db := DbConnect()

	for _, user := range responseObject {

		stmt, err := db.Prepare("INSERT INTO user(first_name, last_name, nick_name, email, password_, dob, image_, about, public) VALUES(?,?,?,?,?,?,?,?,?);")
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Password, user.Dob, user.Image, user.About, 1)
	}
}

// insert mock post data

func InsertMockPostData() {

	// fetch api for mock data

	// iteration for 50 users
	for i := 1; i < 51; i++ {
		var res *http.Response

		res, _ = http.Get("https://63f35a0e864fb1d60014de90.mockapi.io/users/" + strconv.Itoa(i) + "/posts")

		resData, _ := io.ReadAll(res.Body)

		// Unmarshall http response

		var responseObject []Post

		json.Unmarshal(resData, &responseObject)

		// insert into database
		db := DbConnect()

		for _, post := range responseObject {

			stmt, err := db.Prepare("INSERT INTO post(author, message_, image_, created_at, privacy) VALUES(?,?,?,?,?);")
			if err != nil {
				log.Fatal(err)
			}

			defer stmt.Close()

			stmt.Exec(i, post.Message, post.Image, post.CreatedAt, 0)
		}
	}

}
