package log_system

import (
	"time"
)

type Session struct {
	SessionID int       `db:"session_id" json:"session_id"`
	StartTime time.Time `db:"start_time" json:"start_time"`
}
