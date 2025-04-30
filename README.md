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