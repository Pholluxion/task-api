package config

import "github.com/Pholluxion/task-api/internal/utils"

type Config struct {
	DBName string
	Port   string
	utils.JWTUtils
}

func NewConfig() *Config {
	Port := utils.GetString("PORT", "8080")
	DBName := utils.GetString("DB_NAME", "tasks.db")
	JWTSecret := utils.GetString("JWT_SECRET", "your_secret_key")

	return &Config{
		JWTUtils: utils.JWTUtils{
			SecretKey: []byte(JWTSecret),
		},
		DBName: DBName,
		Port:   Port,
	}
}
