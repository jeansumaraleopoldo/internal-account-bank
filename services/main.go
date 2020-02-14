package services

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbUrl] = "localhost"
	conf[dbName] = "bank"
	return conf
}
