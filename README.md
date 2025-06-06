# TidyBeaver - Log Aggregator

TidyBeaver is a modular log aggregation system written in Go. It collects logs from multiple sources (API, database, OS, microservices), processes and stores them in S3, and supports integration with the ELK stack via Filebeat.

## 📦 Project Structure

TidyBeaver/
│
├── bin/ # Binary output folder
├── cmd/ # Entry points
│ ├── api.go
│ ├── elk.go
│ └── main.go
├── internal/
│ ├── aggregator/ # Core log aggregation logic
│ │ ├── aggregator.go
│ │ └── processor.go
│ ├── api/ # REST API handlers
│ │ ├── handler.go
│ │ └── response.go
│ ├── config/ # Application config
│ │ ├── config.go
│ │ └── config.json
│ ├── elk/ # ELK-specific routing and controllers
│ │ ├── controller/
│ │ │ └── base.go
│ │ └── router/
│ │ └── router.go
│ ├── msvc/ # Microservice log logic
│ │ ├── handler.go
│ │ └── response.go
│ ├── sources/ # Source interfaces for logs
│ │ ├── api.go
│ │ ├── db.go
│ │ ├── fs.go
│ │ ├── mock.go
│ │ ├── msvc.go
│ │ └── os.go
│ ├── storage/ # Storage handlers
│ │ ├── db.go
│ │ ├── fs.go
│ │ └── s3/
│ │ ├── createBucket.go
│ │ ├── downloadBucket.go
│ │ ├── listBuckets.go
│ │ ├── s3.go
│ │ └── uploadLogs.go
│ └── logs/ # Log processing (empty or WIP)
│
├── pkg/
│ └── models/ # Domain models for logs
│ ├── aggregatedLogs.go
│ ├── apiLogs.go
│ ├── apiResponse.go
│ ├── dbLogs.go
│ ├── fsLogs.go
│ ├── mockedLogs.go
│ ├── msvcLogs.go
│ └── osLogs.go
│
├── filebeat/ # Filebeat integration
│ ├── Dockerfile
│ └── filebeat.yml
├── scripts/ # SQL scripts
│ ├── create_logs_table.sql
│ ├── get_random_db_events.sql
│ └── insert_logs_table.sql
├── test/ # Unit and integration tests
│ ├── aggregator_test.go
│ └── api_test.go
│
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── tidybeaver-logs/ # Local log output (?)
├── .gitignore
└── README.md # You're here :)


## 🧱 Architecture

+-------------+ +--------------+ +----------------+
| Sources |-----> | Aggregator | ---> | Storage (S3) |
| (API, DB, | | Processor | | or Local FS) |
| OS, MSVC) | +------+-------+ +----------------+
+-------------+ |
v
+---------------+
| Filebeat + ELK|
+---------------+


- **Sources**: Plugins or connectors to log-producing systems.
- **Aggregator**: Gathers and normalizes logs.
- **Processor**: Formats, filters, or enriches logs.
- **Storage**: Stores logs in S3 or a file system.
- **ELK Integration**: Optional pipeline via Filebeat → Logstash → Elasticsearch.

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- Docker + Docker Compose
- AWS CLI (if using S3)
- PostgreSQL or MySQL (for DB logs)

### Setup

```bash
git clone https://github.com/yourname/tidybeaver.git
cd tidybeaver
cp .env.example .env
```

Update .env and config/config.json with your desired environment variables and connection settings.

# Build and run the app
go build -o bin/main ./cmd/main.go
./bin/main

Or use Docker:

docker-compose up --build

Filebeat + ELK (Optional)
Start Filebeat + ELK stack (configured in filebeat/filebeat.yml):

docker-compose -f docker-compose.yml up -d filebeat elk

🧪 Running Tests

go test ./test/...

🛠️ SQL Scripts
create_logs_table.sql - Creates required schema

insert_logs_table.sql - Inserts test logs

get_random_db_events.sql - Simulates log events

📂 Logs
Logs may be output locally in the tidybeaver-logs/ folder or sent to S3.



