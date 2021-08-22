package utils

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DateFormatter(year, month, day, hour, min uint) string {
	return fmt.Sprintf("%04d-%02d/%02d-%02d:%02d", year, month, day, hour, min)
}
