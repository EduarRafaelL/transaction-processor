-- Create base schema
CREATE TABLE clients (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE  transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    client_id VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    transaction_type_id INT NOT NULL REFERENCES transaction_types(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    html_body TEXT NOT NULL
);



-- Insert clients
INSERT INTO clients (id, name, email) VALUES
('12345678', 'Juan Pérez', 'eduar.loopez@gmail.com'),
('87654321', 'Ana Gómez', 'ana.gomez@example.com'),
('11112222', 'Carlos Ramírez', 'carlos.ramirez@example.com');

-- Insert templates
INSERT INTO templates (name, subject, html_body)
VALUES (
    'Stori Transaction Summary Template',
    'Your Monthly Transaction Summary',
    '<p>Hello {{.userName}},</p><p>Here is your transaction summary:</p><ul><li><strong>Total Balance:</strong> ${{.total_balance}}</li><li><strong>Total Transactions:</strong> {{.total_transactions}}</li><li><strong>Total Credit Transactions:</strong> {{.total_credit_transactions}}</li><li><strong>Total Debit Transactions:</strong> {{.total_debit_transactions}}</li><li><strong>Average Credit Amount:</strong> ${{.average_credit_transactions}}</li><li><strong>Average Debit Amount:</strong> ${{.average_debit_transactions}}</li></ul><p>Transactions by Month:</p><ul>{{range $month, $count := .total_transactions_by_month}}<li><strong>{{$month}}:</strong> {{$count}} transaction(s)</li>{{end}}</ul><p>Thank you for being part of Stori.</p>'
);

INSERT INTO templates (name, subject, html_body)
VALUES (
    'Stori Base Template',
    'Your Monthly Transaction Summary',
    '<!DOCTYPE html><html lang="es"><head><meta charset="UTF-8" /><meta http-equiv="X-UA-Compatible" content="IE=edge" /><meta name="viewport" content="width=device-width, initial-scale=1.0" /><title>Correo Stori</title><style>body{font-family:Arial,sans-serif;background-color:#f5f5f5;color:#333;margin:0;padding:0}.container{max-width:600px;margin:auto;background-color:#ffffff;padding:20px;border-radius:8px;box-shadow:0 0 10px rgba(0,0,0,0.1)}.header{background-color:#0051a4;padding:15px;text-align:center;color:white;font-size:24px;font-weight:bold}.header img{max-width:150px;height:auto;display:block;margin:auto}.content{padding:20px;text-align:left;color:#333}.footer{text-align:center;padding:20px;font-size:14px;color:#777}.footer a{color:#0051a4;text-decoration:none;margin:0 10px}.footer img{max-width:20px;vertical-align:middle;margin-right:5px}</style></head><body><div class="container"><div class="header"><img src="https://play-lh.googleusercontent.com/oXTAgpljdbV5LuAOt1NP9_JafUZe9BNl7pwQ01ndl4blYL4N4IQh4-n456P5l_hc1A" alt="Stori Logo" />Stori</div><div class="content">{{.body}}</div><div class="footer"><a href="https://www.storicard.com/"><img src="https://play-lh.googleusercontent.com/oXTAgpljdbV5LuAOt1NP9_JafUZe9BNl7pwQ01ndl4blYL4N4IQh4-n456P5l_hc1A" alt="Web" />https://www.storicard.com</a></div></div></body></html>'
);

INSERT INTO transaction_types (id, name) VALUES
(1, 'credit'),
(2, 'debit');