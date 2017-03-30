package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()

	//Initialize key
	initKeys()

	// open database connection
	createDatabaseConnection()
}

func ShortDown() {
	closeDatabaseConnection()
}
