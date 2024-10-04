# cms-be-service

clone app

run "go mod tidy" or "go mod download" those command for install package

and then create .env file

inside .env file add configuration to connect your local/deploy database and default port for this service.

PORT=8081
<!-- customize to your database configuration-->
DB_HOST="localhost"
DB_NAME="cms_test"
DB_USER="rama"
DB_PASS="xafdewrg"
DB_PORT=5432

<!-- Query add table users in postgres-->
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    email VARCHAR(50),
    password VARCHAR(255),
    role VARCHAR(15),
    image TEXT,
    phone_number VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

run project go run ./cmd/main.go

go to postman or insomnia.
post user data :

post : http://localhost:8081/user

request json :
{
    "first_name": "rama",
    "last_name":"rkt",
    "emai":"rama@mail.com",
    "password":"123456",
    "role":"Admin",
    "image":"rama.pmg",
    "phone_number":"0832425334"
}

<!-- I haven't created a service to view user data so to see if it works look in your database. -->