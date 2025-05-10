package storage

import (
	"database/sql"
	"log"
	"math/rand"
	models "tidybeaver/pkg/models"

	_ "github.com/lib/pq"
)

func WriteSampleLogsToDB(sampleLogs models.SampleLogs) {
	connStr := "user=postgres password=postgres dbname=TidyBeaverLogs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range sampleLogs.SampleLog {
		var userID int
		instanceID := rand.Int63()

		err = db.QueryRow(`INSERT INTO public."Logs" (InstanceID, Time, Level, Source, Service, EntryType, Message, SentToS3)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, instanceID, id.Time, id.Level, "SampleLog", id.Service, "SampleLog", id.Message, "False").Scan(&userID)
		if err != nil {
			log.Println("Error inserting log entry:", err)
			continue
		}

		log.Printf("Log entry inserted with ID: %d\n", instanceID)
	}
}

func WriteLogsToDB(logs any) {
	connStr := "user=postgres dbname=TidyBeaverLogs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var userid int

	err = db.QueryRow(`INSERT INTO Logs (InstanceID, Time, Level, Source, Service, EntryType, Message, SentToS3)
	VALUES (    $age,    NOW(),   'INFO',    'ApplicationLogger',    'AuthService',    'Log',    'User login successful.',    FALSE) RETURNING id`).Scan(&userid)
	if err != nil {
		log.Fatal(err)
	}
}
