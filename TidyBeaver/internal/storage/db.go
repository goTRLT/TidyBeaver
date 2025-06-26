package storage

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strconv"
	models "tidybeaver/pkg/models"
	"time"

	"github.com/lib/pq"
)

func DBStoreLogs(als *models.AggregatedLogs) {
	timeoutSecondsStr := os.Getenv("DB_TIMEOUTSECONDS")
	timeoutSeconds, err := strconv.Atoi(timeoutSecondsStr)
	if err != nil {
		log.Fatal(err)
	}

	connStr := `host=` + os.Getenv("DB_HOST") + ` port=` + os.Getenv("DB_PORT") + ` user=` + os.Getenv("DB_USER") + ` password=` + os.Getenv("DB_PW") + ` dbname=` + os.Getenv("DB_NAME") + ` sslmode=` + os.Getenv("DB_SSLMODE")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds*int(time.Second)))
	defer cancel()
	if ctx == nil {
		log.Println("Error on setting Database context")
	}

	for _, id := range als.AggregatedLog {
		var userID int
		instanceID := rand.Int63()

		err = db.QueryRowContext(ctx, `INSERT INTO public."logs" (category, category_number, checksum, client_ip, "column", component, computer_name, "constraint", container, correlation_id, data, datatype, detail, endpoint, entry_type, environment, errcode, event_id, event_type, file_path, file_size, host, http_method, index, instance_id, latency_ms, level, line_number, log_name, machine_name, message, request_body, replacement_strings, response_body, rows_affected, schema, service, source, split_lines, span_id, status_code, table_name, time_generated, time_written, transaction_id, user_agent, user_id, user_name, query)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49) RETURNING id`,
			id.Category, id.CategoryNumber, id.Checksum, id.ClientIP, id.Column, id.Component, id.ComputerName, id.Constraint, id.Container, id.CorrelationID, pq.Array(id.Data), id.Datatype, id.Detail, id.Endpoint, id.EntryType, id.Environment, id.Errcode, id.EventID, id.EventType, id.Path, id.FileSize, id.Host, id.HTTPMethod, id.Index, id.InstanceID, id.LatencyMs, id.Level, id.LineNumber, id.LogName, id.MachineName, id.Message, id.RequestBody, pq.Array(id.ReplacementStrings), id.ResponseBody, id.RowsAffected, id.Schema, id.Service, id.Source, id.SplitLines, id.SpanID, id.StatusCode, id.TableName, id.TimeGenerated, id.TimeWritten, id.TransactionID, id.UserAgent, id.UserID, id.UserName, id.Query).Scan(&userID)

		if err != nil {
			log.Println("Error inserting log entry:", err)
			continue
		}
		log.Printf("Log entry inserted from Source: %s\n", id.Source)
		log.Printf("Log entry inserted with ID: %d\n", instanceID)
		time.Sleep(200 * time.Millisecond)
	}
}
