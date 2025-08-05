package config

import "os"

// definig types just like we do in the typeScript
type Config struct {
	MongoURI string
	DBName   string
}

/*
This is a function that returns a Config object.
It uses getEnv(...) to try to read environment variables.
If not set, it uses default values.
This helps you configure your app without hardcoding values.
*/
func LoadConfig() Config {
	return Config{
		MongoURI: getEnv("MONGO_URI", "mongodb+srv://jainbhav0207:qXjaBrK6TDukmtJY@cluster0.g5yofar.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"),
		DBName:   getEnv("DB_Name", "EcomUserService"),
	}
}

// this function looks up the environment variable if ant string or any relevent data is available
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
