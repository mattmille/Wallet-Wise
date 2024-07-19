package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"wallet-wise/shared"
	"wallet-wise/src/models"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ctx = context.Background()

// getExpense
func getExpense(c *gin.Context) {
	id := c.Param("id")

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// postExpense adds a charge from JSON received in the request body.
func postExpense(c *gin.Context) {
	var newExpense shared.Expense

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newExpense); err != nil {
		return
	}

	fmt.Println("blah")

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
