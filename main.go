package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wallet-wise/models"
	"wallet-wise/shared"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var router = gin.Default()
var ctx = context.Background()

var DB *sql.DB

// getExpense
func getExpense(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 6, 12)

	expense, err := models.Expenses(qm.Where("expense_id = ?", id)).One(ctx, DB)
	if err != nil {
		log.Fatal("Error: Cannot open table categories: ", err)
	}

	c.Header("Content-Type", "application/json")

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": id,
	// })

	c.JSON(http.StatusOK, expense)
}

// postExpense adds a charge from JSON received in the request body.
func postExpense(c *gin.Context) {
	var newExpense shared.Expense

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newExpense); err != nil {
		return
	}

	var expense models.Expense

	expense.ExpenseID = null.Int64From(newExpense.ID)
	expense.ExpenseAmount = newExpense.Amount
	expense.ExpenseCategoryID = null.Int64From(newExpense.CategoryID)
	expense.ExpenseDescription = newExpense.Description
	expense.ExpenseDate = null.Int64From(newExpense.Date)

	// Fill in the fields of the Expense struct here
	// For example:
	// Amount:      100.50,
	// Description: "Office supplies",
	// Date:        time.Now(),

	err := expense.InsertG(ctx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	// Add the new album to the slice.
	//albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newExpense)
}

func addExpenseRoutes(rg *gin.RouterGroup) {
	expense := rg.Group("/expense")

	expense.GET("/:id", getExpense)
	expense.POST("/", postExpense)
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}

func openDB(dbPath string) (*sql.DB, error) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Successfully connected to the SQLite database")

	return db, nil
}

func main() {
	dbPath := "./db.sqlite"
	var err error

	DB, err = openDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer DB.Close() // Ensure the database connection is closed when the program exits
	boil.SetDB(DB)

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))
	api := router.Group("/api")
	addExpenseRoutes(api)
	v2 := router.Group("/v2")
	addPingRoutes(v2)

	router.Run(":3000")
}
