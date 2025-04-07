package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

// User struct with proper JSON tags
type User struct {
	ID       int    `json:"userid,omitempty"`
	FNAME    string `json:"fname"`
	LNAME    string `json:"lname"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"-"`
	ROLE     string `json:"role"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:103001@tcp(127.0.0.1:3306)/elibrary")
	if err != nil {
		log.Fatal(err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}

// Helper functions
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   message,
	})
}

func sendSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// CREATE USER
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Check if email exists
	var existingID int
	err := db.QueryRow("SELECT userid FROM userdb WHERE email = ?", user.EMAIL).Scan(&existingID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		sendErrorResponse(w, "Database error", http.StatusInternalServerError)
		return
	}
	if existingID > 0 {
		sendErrorResponse(w, "Email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.DefaultCost)
	if err != nil {
		sendErrorResponse(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	result, err := db.Exec("INSERT INTO userdb (fname, lname, email, password, role) VALUES (?, ?, ?, ?, 'user')",
		user.FNAME, user.LNAME, user.EMAIL, hashedPassword)
	if err != nil {
		sendErrorResponse(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	sendSuccessResponse(w, map[string]interface{}{"userid": id})
}

// GET ALL USERS
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT userid, fname, lname, email FROM userdb")
	if err != nil {
		sendErrorResponse(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FNAME, &user.LNAME, &user.EMAIL); err != nil {
			sendErrorResponse(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	sendSuccessResponse(w, users)
}

// UPDATE USER
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Get existing user data
	var existingUser User
	err := db.QueryRow("SELECT password FROM userdb WHERE userid = ?", userID).
		Scan(&existingUser.PASSWORD)
	if err != nil {
		sendErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	// Update password if provided
	password := existingUser.PASSWORD
	if user.PASSWORD != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.DefaultCost)
		if err != nil {
			sendErrorResponse(w, "Error processing request", http.StatusInternalServerError)
			return
		}
		password = string(hashedPassword)
	}

	_, err = db.Exec("UPDATE userdb SET fname = ?, lname = ?, email = ?, password = ? WHERE userid = ?",
		user.FNAME, user.LNAME, user.EMAIL, password, userID)
	if err != nil {
		sendErrorResponse(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	sendSuccessResponse(w, "User updated")
}

// DELETE USER
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	_, err := db.Exec("DELETE FROM userdb WHERE userid = ?", userID)
	if err != nil {
		sendErrorResponse(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	sendSuccessResponse(w, "User deleted")
}

// GET USER BY ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user User
	err := db.QueryRow("SELECT userid, fname, lname, email FROM userdb WHERE userid = ?", userID).
		Scan(&user.ID, &user.FNAME, &user.LNAME, &user.EMAIL)
	if err != nil {
		sendErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	sendSuccessResponse(w, user)
}

// LOGIN USER
func loginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		EMAIL    string `json:"email"`
		PASSWORD string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		sendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var dbUser User
	err := db.QueryRow("SELECT userid, fname, lname, email, password, role FROM userdb WHERE email = ?", credentials.EMAIL).
		Scan(&dbUser.ID, &dbUser.FNAME, &dbUser.LNAME, &dbUser.EMAIL, &dbUser.PASSWORD, &dbUser.ROLE)
	if err != nil {
		sendErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.PASSWORD), []byte(credentials.PASSWORD)); err != nil {
		sendErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	responseUser := User{
		ID:    dbUser.ID,
		FNAME: dbUser.FNAME,
		LNAME: dbUser.LNAME,
		EMAIL: dbUser.EMAIL,
		ROLE: dbUser.ROLE,
	}
	sendSuccessResponse(w, responseUser)
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()

	// Configure routes
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", getUserByID).Methods("GET")
	r.HandleFunc("/login", loginUser).Methods("POST")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server started at :8080")
	log.Fatal(server.ListenAndServe())
}