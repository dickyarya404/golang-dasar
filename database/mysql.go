package database

var connection string

func init() {
	connection = "MySQLConnection"
}

func GetDatabase() string {
	return connection
}