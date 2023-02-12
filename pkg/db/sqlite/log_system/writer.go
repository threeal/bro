package log_system

import "database/sql"

type DBWriter struct {
	db        *sql.DB
	sessionID int
}

func (w *DBWriter) Write(p []byte) (n int, err error) {
	QueryDB(w.db, InsertToLogsTableSQL, nil, string(p), w.sessionID)
	if err != nil {
		return
	}
	return len(p), nil
}

func (w *DBWriter) SetDB(db *sql.DB) {
	w.db = db
}

func (w *DBWriter) SetSessionID(sessionID int) {
	w.sessionID = sessionID
}
