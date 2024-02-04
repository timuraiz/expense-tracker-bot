package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timuraiz/expense-tracker-bot/pkg/config"
	"time"
)

type User interface {
	Authenticate(userID, password string) (bool, error)
}

type Expense interface {
	AddExpense(expense ExpenseDetail) (string, error) // Returns ExpenseID on success
	EditExpense(expenseID string, expense ExpenseDetail) error
	DeleteExpense(expenseID string) error
}

type ExpenseDetail struct {
	UserID      string
	Amount      float64
	Category    string
	Date        time.Time
	Currency    string
	Description string
	Tags        []string
}

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open(cfg.Driver, sqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
