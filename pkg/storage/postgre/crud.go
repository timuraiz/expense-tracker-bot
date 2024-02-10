package postgre

import (
	"database/sql"
	"fmt"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type PostgresRepository struct {
	db      *sql.DB
	queries SQLQueries
}
type SQLQueries struct {
	CreateTables  []string `yaml:"create_tables"`
	AddExpense    string   `yaml:"add_expense"`
	EditExpense   string   `yaml:"edit_expense"`
	DeleteExpense string   `yaml:"delete_expense"`
}

func loadSQLQueries(filePath string) (SQLQueries, error) {
	var queries SQLQueries
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return queries, err
	}

	err = yaml.Unmarshal(content, &queries)
	return queries, err
}

func NewPostgresRepository(db *sql.DB) (*PostgresRepository, error) {
	queries, err := loadSQLQueries("pkg/storage/queries.yaml")
	if err != nil {
		return nil, err
	}

	repo := &PostgresRepository{db: db, queries: queries}
	err = repo.initTables()
	return repo, err
}
func (pr *PostgresRepository) initTables() error {
	for _, query := range pr.queries.CreateTables {
		if _, err := pr.db.Exec(query); err != nil {
			return err
		}
	}
	return nil
}

func (pr *PostgresRepository) AddExpense(expense storage.ExpenseDetail) (int, error) {
	var insertedId int

	err := pr.db.QueryRow(pr.queries.AddExpense, expense.UserID, expense.Amount, expense.Category, expense.Description, expense.Date).Scan(&insertedId)
	if err != nil {
		return -1, fmt.Errorf("error adding expense: %w", err)
	}

	return insertedId, nil
}

func (pr *PostgresRepository) EditExpense(expenseID string, expense storage.ExpenseDetail) error {
	_, err := pr.db.Exec(pr.queries.EditExpense, expense.Amount, expense.Category, expense.Description, expense.Date, expenseID)
	if err != nil {
		return fmt.Errorf("error updating expense: %w", err)
	}

	return nil
}

func (pr *PostgresRepository) DeleteExpense(expenseID string) error {
	_, err := pr.db.Exec(pr.queries.DeleteExpense, expenseID)
	if err != nil {
		return fmt.Errorf("error deleting expense: %w", err)
	}

	return nil
}
