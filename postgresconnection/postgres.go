package postgresconnection

import (
	"github.com/jackc/pgx/v5"
	"os"
	"log"
	"context"
)

func GetConnection() *pgx.Conn {
	psql_user := os.Getenv("PSQL_USER")
	psql_pass := os.Getenv("PSQL_PASS")
	psql_host := os.Getenv("PSQL_HOST")
	psql_port := os.Getenv("PSQL_PORT")
	psql_db := os.Getenv("PSQL_DB")
	conn, err := pgx.Connect(context.Background(), "postgresql://"+psql_user+":"+psql_pass+"@"+psql_host+":"+psql_port+"/"+psql_db)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}