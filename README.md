# GIN Restfull API
### Simple Book Library Borrowing API ###
`` How to Run ``
- Clone the project 
- Create a .env file and fill the database details that can be find in the .env-example file
- Create Tables and Populate Data using the commands on the ddl-dml.sql file
- Run the project
- Create server.pem(public) and key.pem(private) using TLS RSA key with this command `openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes`