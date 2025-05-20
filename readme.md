# ddp

#### Docker Usage
Ensure you have docker + docker compose on your machine.  
Then either use docker desktop (I don't know how to use it so good luck)
or the docker cli

run:
```sh
docker compose up --build # might need 'sudo' on mac/linux depending on setup
```
then wait for it to say `Server listening http://127.0.0.1:8000`

#####
Database Details:  
```
Login: root : 1234  # admin user
       webapp_user : WebAppDB2025! # restricted user
```
Website Details:
```
Club Recorder Account: AA12345 : password
Normal User: AA12346 : password
```
#### Manual Usage 
Expects a db to be running on localhost:3306
with a user root, password 1234

change this in web/main.go `dsn`

requires [go](https://go.dev/doc/install)

```sh
cd web/
go run .
```
[https://localhost:8000](https://localhost:8000)
