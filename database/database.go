package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
	"sync"
	"time"

	fileops "github.com/babbage88/go-history/utils/fileops"

	_ "github.com/mattn/go-sqlite3"
)

type CommandHistoryDbEntry struct {
	Id           int64     `json:"id"`
	DateExecuted time.Time `json:"date_executed"`
	BaseCommand  string    `json:"base_ommand"`
	SubCommand   []string  `json:"sub_command"`
	LineNumber   int64     `json:"line_number"`
}

type DatabaseConnection struct {
	DbName string `json:"database"`
	DbPath string `json:"dbPath"`
}

type DatabaseConnectionOptions func(*DatabaseConnection)

// Global db instance
var (
	db     *sql.DB
	dbOnce sync.Once
	dbErr  error
)

func InitializeDbConnection(dbConn *DatabaseConnection) (*sql.DB, error) {
	dbOnce.Do(func() {
		// Connect to the SQLite database
		slog.Info("Connecting to database: " + dbConn.DbName)

		// SQLite uses a file-based connection string
		dbFilePath := filepath.Join(dbConn.DbPath, dbConn.DbName)
		dsn := fmt.Sprintf("%s", dbFilePath)
		db, dbErr = sql.Open("sqlite3", dsn)
		if dbErr != nil {
			slog.Error("Error connecting to the database", slog.String("Error", dbErr.Error()))
			return
		}

		// Create the table if it doesn't exist
		dbErr = createTable(db)
		if dbErr != nil {
			slog.Error("Error creating table", slog.String("Error", dbErr.Error()))
		}
	})
	return db, dbErr
}

func createTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS commands (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date_executed TEXT,
		base_command TEXT,
		sub_command TEXT,
		line_number INTEGER
	);`

	slog.Info("Creating table: commands")

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	slog.Info("Table created successfully")
	return nil
}

func CloseDbConnection() error {
	if db != nil {
		slog.Info("Closing DB Connection")
		return db.Close()
	}
	return nil
}

func NewDatabaseConnection(opts ...DatabaseConnectionOptions) *DatabaseConnection {
	defaultDbPath, _ := fileops.GetCurrentUserPath()

	dbName := "dev-bash-history.db"

	dbPath := defaultDbPath

	db := &DatabaseConnection{
		DbName: dbName,
		DbPath: dbPath,
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

func WithDbName(DbName string) DatabaseConnectionOptions {
	return func(c *DatabaseConnection) {
		c.DbName = DbName
	}
}

func WithDbPath(DbPath string) DatabaseConnectionOptions {
	return func(c *DatabaseConnection) {
		c.DbPath = DbPath
	}
}

func InsertCommandHistoryEntries(entries []fileops.CommandHistoryEntry) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	insertSQL := `INSERT INTO commands (date_executed, base_command, sub_command, line_number) VALUES (?, ?, ?, ?)`

	for _, entry := range entries {
		subCommand := strings.Join(entry.SubCommand, " ")
		_, err := db.Exec(insertSQL, entry.DateExecuted.Format(time.RFC3339), entry.BaseCommand, subCommand, entry.LineNumber)
		if err != nil {
			slog.Error("Error inserting command history entry", slog.String("Error", err.Error()))
			return err
		}
	}

	slog.Info("Command history entries inserted successfully")
	return nil
}

func GetAllCmdHistory(db *sql.DB) ([]CommandHistoryDbEntry, error) {
	query := `SELECT id, date_executed, base_command, sub_command, line_number FROM commands`

	rows, err := db.Query(query)
	if err != nil {
		slog.Error("Error running query", slog.String("Error", err.Error()))
	}

	var commands []CommandHistoryDbEntry
	for rows.Next() {
		var command CommandHistoryDbEntry

		if err := rows.Scan(&command.Id,
			&command.DateExecuted,
			&command.BaseCommand,
			&command.SubCommand,
			&command.LineNumber); err != nil {
			slog.Error("Error parsing DB response", slog.String("Error", err.Error()))
		}
		slog.Info("Appending DNS Record", slog.String("id", fmt.Sprint(command.Id)), slog.String("base_command:", command.BaseCommand))
		commands = append(commands, command)
	}

	return commands, nil

}
