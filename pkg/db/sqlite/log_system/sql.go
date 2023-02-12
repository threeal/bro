package log_system

const (
	CreateLogsTableSQL = `
	CREATE TABLE IF NOT EXISTS logs (
		log_id INTEGER PRIMARY KEY,
		topic VARCHAR(64),
		content TEXT NOT NULL,
		time DATETIME DEFAULT CURRENT_TIMESTAMP,
		session_id INTEGER NOT NULL,
		FOREIGN KEY (session_id) 
		REFERENCES sessions (session_id) 
			ON UPDATE CASCADE 
			ON DELETE CASCADE
	);
`
)

const (
	EnableForeignKey = `
	PRAGMA foreign_keys = ON;
`
)

const (
	CreateSessionsTableSQL = `
	CREATE TABLE IF NOT EXISTS sessions (
		session_id INTEGER PRIMARY KEY,
		start_time DATETIME DEFAULT CURRENT_TIMESTAMP
	);
`
)

const (
	InsertToLogsTableSQL = `
	INSERT INTO logs(topic,content,session_id) VALUES(?,?,?);
`
)

const (
	InsertToSessionsTableSQL = `
	INSERT INTO sessions DEFAULT VALUES;
`
)
