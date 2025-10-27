package config

import "github.com/Pholluxion/task-api/internal/utils"

type Config struct {
	DBName          string
	Port            string
	SecretKey       string
	TokenExpireTime int
}

func NewConfig() *Config {
	Port := utils.GetString("PORT", "8080")
	DBName := utils.GetString("DB_NAME", "tasks.db")
	JWTSecret := utils.GetString("JWT_SECRET", "your_secret_key")
	TokenExpireTime := utils.GetInt("TOKEN_EXPIRE_TIME", 30)

	return &Config{
		SecretKey:       JWTSecret,
		DBName:          DBName,
		Port:            Port,
		TokenExpireTime: TokenExpireTime,
	}
}
