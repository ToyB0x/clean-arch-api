package config

var Configs *configs

type configs struct {
	APP_ENV            string
	MIGRATION_FILE_DIR string

	DB_NAME          string
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASS          string
	DB_INSTANCE_NAME string

	REDIS_PORT int

	MaxRequestSize int64
}
