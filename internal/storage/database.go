package storage

import (
	"database/sql"
	"log"
	"math/rand"
	models "tidybeaver/pkg/models"

	_ "github.com/lib/pq"
)

func WriteSampleLogsToDB(sampleLogs models.SampleLogs) {
	connStr := "host=localhost port=5432 user=tidybeaver password=tidybeaver dbname=TidyBeaverLogs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range sampleLogs.SampleLog {
		var userID int
		instanceID := rand.Int63()

		err = db.QueryRow(`INSERT INTO public."Logs" (instanceID, time, level, source, service, entrytype, message, awsbucketsent)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, instanceID, id.Time, id.Level, "SampleLog", id.Service, "SampleLog", id.Message, "False").Scan(&userID)
		if err != nil {
			log.Println("Error inserting log entry:", err)
			continue
		}

		log.Printf("Log entry inserted with ID: %d\n", instanceID)
	}
}

func WriteLogsToDB(logs models.TransformedLogs) {
	connStr := "host=localhost port=5432 user=tidybeaver password=tidybeaver dbname=TidyBeaverLogs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range logs.TransformedLog {
		var userID int
		instanceID := rand.Int63()

		err = db.QueryRow(`INSERT INTO public."Logs" (instanceID, time, level, source, service, entrytype, message, awsbucketsent)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, instanceID, id.Time, id.Level, "SampleLog", id.Service, "SampleLog", id.Message, "False").Scan(&userID)
		if err != nil {
			log.Println("Error inserting log entry:", err)
			continue
		}

		log.Printf("Log entry inserted with ID: %d\n", instanceID)
	}
}
