# Transaction Processor and Email Sender

## Overview

This repository contains two services: `processor` and `email_sender`.

- The `processor` service includes a file watcher that detects new `.csv` files in the directory specified by the `INPUT_PATH` environment variable. When a new file is detected, it extracts the user ID from the filename, processes the contained transactions, and calculates summary statistics such as:

  - Number of credit and debit transactions
  - Average amounts per transaction type
  - Total account balance
  - Number of transactions per month

  After processing, the data is saved in a PostgreSQL database as a transaction history and a summary is sent to the `email_sender` service.

- The `email_sender` service is an HTTP API that receives transaction summaries, fetches the corresponding templates from the database, renders them, and sends the final email via `SMTP` based on the environment configuration.

---

## System Architecture

- **Processor:** Monitors the input folder, processes CSV transaction files, stores data in the database, and sends summaries via HTTP.
- **Email Sender:** HTTP service that composes and sends emails using SMTP.
- **PostgreSQL:** Stores clients, transaction data, and HTML email templates.

---

## Technologies Used

- Golang
- PostgreSQL
- fsnotify
- net/smtp
- Docker
- Docker Compose

---

## Prerequisites

- Docker and Docker Compose installed ([Official Docker guide](https://docs.docker.com/get-docker/))
- Properly configured environment variables for both services
- SQL scripts executed to initialize the required database tables and insert default templates
- A working SMTP account (e.g., Gmail) with an application-specific password ([How to generate an app password](https://support.google.com/accounts/answer/185833))

---

## Environment Variables

### Email Sender

```bash
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password
EMAIL_BODY_TEMPLATE="Exact name of the body template"
EMAIL_MESSAGE_TEMPLATE="Exact name of the message template"
```

### PostgreSQL
```bash
POSTGRES_USER=storiuser
POSTGRES_PASSWORD=storipassword
POSTGRES_DB=storidb
```

### Processor
```bash
INPUT_PATH=/input
OUTPUT_PATH=/output
DELIMITER=,
DB_HOST=postgres
DB_PORT=5432
DB_USER=storiuser
DB_PASSWORD=storipassword
DB_NAME=storidb
EMAIL_SERVICE_URL=http://email-sender:8080/process-request-email
```

### SQL Scripts
SQL scripts for creating the necessary tables (clients, transactions, templates) and sample data inserts are located in the /scripts folder.

### Usage Instructions
Clone the repository.

Configure the required environment variables.

Run the SQL scripts to set up the database schema and insert initial data.

Start the services using:

```bash
docker-compose up --build
```

Add a .csv file to the input/ directory, named with a valid user ID (e.g., 12345678.csv).

The processor service will automatically detect the file, process it, and trigger the email summary to be sent.

Sample CSV File
```csv
id,date,amount
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10
```

### Error Handling
If a file fails to process, an error log will be generated in the output/ directory with the filename and error details.

The system remains active and continues processing other files.

### Technical Considerations
The system uses fsnotify to simulate AWS Lambda/S3 triggers by detecting new files in real time without polling.

Email templates are stored in the database and rendered dynamically at runtime.