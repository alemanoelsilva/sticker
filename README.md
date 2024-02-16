# Sticker

A Go lang backend application to make commitments and stick to it.

## Basic Configuration

Create a `.env` file with

- DB_STRING: a connection mysql string as `user:pas@tcp(127.0.0.1:3306)/db`
- PORT: application HTTP port
- JWT_TOKEN: jwt token
- MYSQL_ROOT_PASSWORD: db root password
- MYSQL_DATABASE: database name
- MYSQL_USER: user name
- MYSQL_PASSWORD: user password

## Running

To run locally run `docker-compose up --build`

## Testing

[TODO]