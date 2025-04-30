
INSERT INTO templates (name, subject, html_body)
VALUES (
    'Stori Transaction Summary Template',
    'Your Monthly Transaction Summary',
    '<p>Hello {{.user_name}},</p><p>Here is your transaction summary:</p><ul><li><strong>Total Balance:</strong> ${{.total_balance}}</li><li><strong>Total Transactions:</strong> {{.total_transactions}}</li><li><strong>Total Credit Transactions:</strong> {{.total_credit_transactions}}</li><li><strong>Total Debit Transactions:</strong> {{.total_debit_transactions}}</li><li><strong>Average Credit Amount:</strong> ${{.average_credit_transactions}}</li><li><strong>Average Debit Amount:</strong> ${{.average_debit_transactions}}</li></ul><p>Transactions by Month:</p><ul>{{range $month, $count := .total_transactions_by_month}}<li><strong>{{$month}}:</strong> {{$count}} transaction(s)</li>{{end}}</ul><p>Thank you for being part of Stori.</p>'
);

INSERT INTO templates (name, subject, html_body)
VALUES (
    'Stori Base Template',
    'Your Monthly Transaction Summary',
    '<!DOCTYPE html><html lang="es"><head><meta charset="UTF-8" /><meta http-equiv="X-UA-Compatible" content="IE=edge" /><meta name="viewport" content="width=device-width, initial-scale=1.0" /><title>Correo Stori</title><style>body{font-family:Arial,sans-serif;background-color:#333;color:#fff;margin:0;padding:0}.container{max-width:600px;margin:auto;background-color:#1a1a1a;padding:15px;border-radius:10px;overflow:hidden}.header{background-color:#0051a4;padding:0;text-align:center}.header img{width:100%;height:auto;display:block}.content{padding:20px;text-align:left;background-color:#1a1a1a;border-top:4px solid #0051a4}.footer{text-align:center;padding:20px;background-color:#1a1a1a}.footer a{color:#fff;text-decoration:none;margin:0 10px;font-size:14px}.footer img{max-width:20px;vertical-align:middle;margin-right:5px}</style></head><body><div class="container"><div class="header"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/30/Stori_logo.svg/2560px-Stori_logo.svg.png" alt="Stori Header" /></div><div class="content"><p>{{.body}}</p></div><div class="footer"><a href="https://www.storicard.com/"><img src="https://play-lh.googleusercontent.com/oXTAgpljdbV5LuAOt1NP9_JafUZe9BNl7pwQ01ndl4blYL4N4IQh4-n456P5l_hc1A" alt="Web" />https://www.storicard.com</a></div></div></body></html>'
);


INSERT INTO clients (id, name, email) VALUES
('12345678', 'Juan Pérez', 'juan.perez@example.com'),
('87654321', 'Ana Gómez', 'ana.gomez@example.com'),
('11112222', 'Carlos Ramírez', 'carlos.ramirez@example.com');

INSERT INTO transaction_types (id, name) VALUES
(1, 'credit'),
(2, 'debit');