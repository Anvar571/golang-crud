package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/Anvar571/golang-crud/configs"
)

func ConnectionDb(cfg *configs.Config) (*sql.DB, error) {
	config := fmt.Sprintf(
		"host=%s port=%s user=%s dbName=%s password=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDb,
		cfg.PostgresPassword,
	)

	sqlDb, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}

	return sqlDb, nil
}
