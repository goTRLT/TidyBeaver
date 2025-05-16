package sources

import (
	"database/sql"
	"log"
	config "tidybeaver/internal/config"
	"tidybeaver/pkg/models"
)

func FetchDBLogs() (models.DBLogs, error) {
	connStr := `host=` + config.ConfigValues.Database.Host + ` port=` + config.ConfigValues.Database.Port + ` user=` + config.ConfigValues.Database.User + ` password=` + config.ConfigValues.Database.Password + ` dbname=` + config.ConfigValues.Database.Name + ` sslmode=` + config.ConfigValues.Database.SSLMode
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM get_random_db_events($1);`, config.LogAmountSet)
	if err != nil {
		log.Println("Error fetching random db events:", err)
	}

	defer rows.Close()

	return rows, err
}
