# Transaction Processor and Email Sender

## Descripción general

Este repositorio contiene dos servicios: `processor` y `email_sender`.

- El servicio `processor` incluye un file watcher que detecta nuevos archivos `.csv` en la carpeta especificada por la variable de entorno `INPUT_PATH`. Al detectar un archivo nuevo, extrae el ID del usuario desde el nombre del archivo, procesa las transacciones contenidas, calcula estadísticas como:

  - Número de transacciones de crédito y débito
  - Promedio de transacciones por tipo
  - Balance total
  - Total de transacciones por mes

  Después guarda los datos en una base de datos PostgreSQL como histórico y envía el resumen al servicio `email_sender`.

- El servicio `email_sender` es una API HTTP que recibe solicitudes con resúmenes de transacciones, obtiene las plantillas desde la base de datos y genera el correo usando `SMTP` en base a la configuración del entorno.

---

## Arquitectura general

- **Processor:** Lee y procesa archivos `.csv` con transacciones, guarda en base de datos y envía resumen por HTTP.
- **Email sender:** Servicio HTTP que arma correos y los envía vía `SMTP`.
- **PostgreSQL:** Almacena clientes, transacciones y plantillas HTML de correos.

---

## Tecnologías utilizadas

- Golang
- PostgreSQL
- fsnotify
- net/smtp
- Docker
- Docker Compose

---

## Requisitos de ejecución

- Tener Docker y Docker Compose instalados ([Guía oficial](https://docs.docker.com/get-docker/)).
- Configurar las variables de entorno para cada servicio.
- Ejecutar los scripts SQL para crear las tablas e insertar plantillas.
- Tener una cuenta SMTP real (como Gmail) con contraseña de aplicación ([cómo generar una app password](https://support.google.com/accounts/answer/185833)).

---

## Variables de entorno

### Email Sender
```bash
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=correo@gmail.com
SMTP_PASSWORD=tu_app_password
EMAIL_BODY_TEMPLATE="Nombre exacto del template de body"
EMAIL_MESSAGE_TEMPLATE="Nombre exacto del template de mensaje"
```
PostgreSQL
```bash
POSTGRES_USER=storiuser
POSTGRES_PASSWORD=storipassword
POSTGRES_DB=storidb
```
Processor
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
##  Scripts SQL
Los scripts para crear las tablas (clients, transactions, templates) y los inserts de ejemplo se encuentran en la carpeta /scripts.

## Instrucciones de uso
Clona el repositorio.

Configura las variables de entorno.

Inserta los datos necesarios en la base de datos usando los scripts SQL.

Ejecuta los servicios con:

```bash
docker-compose up --build
```

Ejemplo de archivo CSV
csv
Copiar
Editar
```csv
id,date,amount
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10
```