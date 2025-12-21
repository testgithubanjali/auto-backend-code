package config

import "os"

type Config struct {
	MongoURI  string
	DBName    string
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:    getEnv("DB_NAME", "auto_booking"),
		JWTSecret: getEnv("JWT_SECRET", "mysecretkey"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
