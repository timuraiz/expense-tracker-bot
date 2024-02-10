package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timuraiz/expense-tracker-bot/pkg/config"
	"time"
)

type Crud interface {
	AddExpense(expense ExpenseDetail) (int, error) // Returns ExpenseID on success
	EditExpense(expenseID string, expense ExpenseDetail) error
	DeleteExpense(expenseID string) error
}

type ExpenseDetail struct {
	UserID      int64
	Amount      float64
	Category    string
	Date        time.Time
	Description string
}

func SetupDatabase(cfg *config.Config) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Open the connection
	db, err := sql.Open(cfg.Driver, connStr)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
