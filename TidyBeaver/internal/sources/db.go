package sources

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
)

var dbLogEntry models.DBLog

func GetDBLogs() (dbLogs *models.DBLogs, err error) {
	dbLogs = &models.DBLogs{}
	logAmount, _ := strconv.ParseInt(config.CFG.App.LogAmount, 0, 0)

	connStr := `host=` + os.Getenv("DB_HOST") + ` port=` + os.Getenv("DB_PORT") + ` user=` + os.Getenv("DB_USER") + ` password=` + os.Getenv("DB_PW") + ` dbname=` + os.Getenv("DB_NAME") + ` sslmode=` + os.Getenv("DB_SSLMODE")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		return dbLogs, err

	}

	rows, err := db.Query(`SELECT * FROM get_random_db_events($1);`, int(logAmount))
	if err != nil {
		log.Println("Error fetching random db events:", err)
		return dbLogs, err
	}

	defer rows.Close()

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
			log.Println(err)
			continue
		}

		dbLogs.DBLog = append(dbLogs.DBLog, dbLogEntry)
	}
	return dbLogs, rows.Err()
}
