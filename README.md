# product-crud

You will need Go, Docker and Docker Compose installed on your computer to run this project.

### Run application:
```sh
docker compose up --build
```

### Secondly, run the migrations to create tables:
```sh
docker compose --profile tools run --rm migrate up
```

You can also drop all the tables using migrate down:
```sh
docker compose --profile tools run --rm migrate down
```

Additionally, to verify that migrations work, you can connect to database through client like PGAdmin or DBeaver or just run the following command and then enter \d (to quit use \q):
```sh
docker compose exec db psql -U postgres -d product-db
```

### To stop the container run:
```sh
docker compose down
```

Note: if you want to make data persistent, just uncomment comments in docker-compose.yml, so that it will create a volume next time you build the project.

## Test API

### Test buyer CRUD
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "contact":"john.doe@gmail.com"}' http://localhost:8090/buyers
curl -X GET http://localhost:8090/buyers
curl -X GET http://localhost:8090/buyers/get?id=1
curl -X PUT -H "Content-Type: application/json" -d '{"id":"1", "name":"Johnson Donovan", "contact":"john.donovan@yahoo.com"}' http://localhost:8090/buyers
curl -X DELETE http://localhost:8090/buyers?id=1
```

### Test purchase CRUD
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"Football Jersey", "description":"Real Madrid season 2022/2023 jersey", "quantity":"1", "price":"50", "buyer_id":"1"}' http://localhost:8090/purchases
curl -X GET http://localhost:8090/purchases
curl -X GET http://localhost:8090/purchases/get?id=1
curl -X PUT -H "Content-Type: application/json" -d '{"id":"1", "name":"Real Updated Jersey", "description":"Real Madrid season 2023/2024 home jersey", "quantity":"3", "price":"75", "buyer_id":"1"}' http://localhost:8090/purchases
curl -X DELETE http://localhost:8090/purchases?id=1
```

## to run go tests run the following:
```sh
go test internal/handler/*.go -v
```