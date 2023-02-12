package log_system

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"

	"github.com/threeal/bro/pkg/utils"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetLogPath() (string, error) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(configDir, "log.db")
	return configPath, nil
}

func CreateDatabaseFile() string {
	logPath, err := GetLogPath()
	if err != nil {
		log.Fatal(err)
	}
	if fileExists(logPath) {
		return logPath
	}
	file, err := os.Create(logPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	return logPath
}

func InitDB() *sql.DB {
	logPath := CreateDatabaseFile()
	db, err := sql.Open("sqlite3", logPath)
	if err != nil {
		fmt.Println(err)
	}
	QueryDB(db, EnableForeignKey)
	QueryDB(db, CreateSessionsTableSQL)
	QueryDB(db, CreateLogsTableSQL)
	return db
}

func QueryDB(db *sql.DB, sqlString string, sqlQuery ...interface{}) {
	log_table := sqlString
	query, err := db.Prepare(log_table)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(sqlQuery...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Query ran successfully!")
}

func QueryLastID(db *sql.DB) (int, error) {
	var lastID int
	// Query for a value based on a single row.
	if err := db.QueryRow("SELECT last_insert_rowid()").Scan(&lastID); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("unknown row")
		}
		return 0, fmt.Errorf("Error: %v", err)
	}
	return lastID, nil
}
