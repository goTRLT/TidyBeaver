package sources

import (
	"database/sql"
	"log"
	"os"
	config "tidybeaver/internal/config"
	"tidybeaver/pkg/models"
)

var dbLogEntry models.DBLog

func FetchDBLogs() (models.DBLogs, error) {
	connStr := `host=` + os.Getenv("DB_HOST") + ` port=` + os.Getenv("DB_PORT") + ` user=` + os.Getenv("DB_USER") + ` password=` + os.Getenv("DB_PW") + ` dbname=` + os.Getenv("DB_NAME") + ` sslmode=` + os.Getenv("SSLMODE")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM get_random_db_events($1);`, config.LogAmountSet)
	if err != nil {
		log.Println("Error fetching random db events:", err)
	}

	defer rows.Close()
	var dbLogs models.DBLogs

	for rows.Next() {
		err := rows.Scan(
			&dbLogEntry.Level,
			&dbLogEntry.Column,
			&dbLogEntry.Constraint,
			&dbLogEntry.Datatype,
			&dbLogEntry.Table_name,
			&dbLogEntry.Schema,
			&dbLogEntry.Errcode,
			&dbLogEntry.Detail,
		)

		if err != nil {
			log.Fatal(err)
			continue
		}

		dbLogs.DBLog = append(dbLogs.DBLog, dbLogEntry)
	}
	// fmt.Println("DbLogs", dbLogs)
	return dbLogs, rows.Err()
}
