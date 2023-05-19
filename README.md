# product-crud

You will need Go, Docker and Docker Compose installed on your computer to run this project.
Application can be easily started using go run command:
```sh
go run cmd/app/main.go
```

However, it cannot work correclty without a database, so you can run it using Docker Compose.

## Firstly, run the database using docker compose:
```sh
docker compose up --build
```

## Secondly, run the migrations to create tables:
```sh
docker compose --profile tools run --rm migrate up
```

## You can drop all the tables using migrate down:
```sh
docker compose --profile tools run --rm migrate down
```

Additionally, to verify that migrations work, you can connect to database through client like PGAdmin or DBeaver or just run the following command and then enter \d (to quit use \q):
```sh
docker compose exec db psql -U postgres -d your_db_name
```

## To stop the container run:
```sh
docker compose down
```

Note: if you want to make data persistent, just uncomment comments in docker-compose.yml, so that it will create a volume next time you build the project.

## Test API

### Test buyer CRUD
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "contact":"john.doe@gmail.com"}' http://localhost:8090/buyer/create
curl -X GET -H http://localhost:8090/buyer/get?id=1
curl -X PUT -H "Content-Type: application/json" -d '{"id":"1", "name":"Johnson Donovan", "contact":"john.donovan@yahoo.com"}' http://localhost:8090/buyer/update
curl -X DELETE -H http://localhost:8090/buyer/delete?id=1
```

### Test purchase CRUD
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "contact":"john.doe@gmail.com"}' http://localhost:8090/purchase/create
curl -X GET -H http://localhost:8090/purchase/get?id=1
curl -X PUT -H "Content-Type: application/json" -d '{"id":"1", "name":"Johnson Donovan", "contact":"john.donovan@yahoo.com"}' http://localhost:8090/purchase/update
curl -X DELETE -H http://localhost:8090/purchase/delete?id=1
```