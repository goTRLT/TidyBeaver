package sources

import (
	"database/sql"
	"log"
	config "tidybeaver/internal/config"
	"tidybeaver/pkg/models"
)

var dbLogEntry models.DBLog

func FetchDBLogs() (models.DBLogs, error) {
	connStr := `host=` + config.EnvVar["DB_HOST"] + ` port=` + config.EnvVar["DB_PORT"] + ` user=` + config.EnvVar["DB_USER"] + ` password=` + config.EnvVar["DB_PW"] + ` dbname=` + config.EnvVar["DB_NAME"] + ` sslmode=` + config.EnvVar["SSLMODE"]
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
