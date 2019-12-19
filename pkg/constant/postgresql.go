package constant

import "os"

var POSTGRESQL_DB_NAME = "agile-journey-demo"
var POSTGRESQL_USER = "postgres"
var POSTGRESQL_PASSWORD = "abc123456"
var POSTGRESQL_HOST = "127.0.0.1"
var POSTGRESQL_PORT = "32432"
var DISABLE_CREATE_TABLE = false
var DISABLE_SQL_GENERATE_PRINT = false

func init() {
	if os.Getenv("POSTGRESQL_DB_NAME") != "" {
		POSTGRESQL_DB_NAME = os.Getenv("POSTGRESQL_DB_NAME")
	}
	if os.Getenv("POSTGRESQL_USER") != "" {
		POSTGRESQL_USER = os.Getenv("POSTGRESQL_USER")
	}
	if os.Getenv("POSTGRESQL_PASSWORD") != "" {
		POSTGRESQL_PASSWORD = os.Getenv("POSTGRESQL_PASSWORD")
	}
	if os.Getenv("POSTGRESQL_HOST") != "" {
		POSTGRESQL_HOST = os.Getenv("POSTGRESQL_HOST")
	}
	if os.Getenv("POSTGRESQL_PORT") != "" {
		POSTGRESQL_PORT = os.Getenv("POSTGRESQL_PORT")
	}
	if os.Getenv("DISABLE_CREATE_TABLE") != "" {
		DISABLE_CREATE_TABLE = true
	}
	if os.Getenv("DISABLE_SQL_GENERATE_PRINT") != "" {
		DISABLE_SQL_GENERATE_PRINT = true
	}
}
