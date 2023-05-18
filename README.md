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