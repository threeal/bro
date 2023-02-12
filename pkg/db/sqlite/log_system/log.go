package log_system

import (
	"database/sql"
	"time"
)

type Log struct {
	LogID     int            `db:"log_id" json:"log_id"`
	Time      time.Time      `db:"time" json:"time"`
	SessionID int            `db:"session_id" json:"session_id"`
	Topic     sql.NullString `db:"topic" json:"topic"`
	Content   string         `db:"content" json:"content"`
}
