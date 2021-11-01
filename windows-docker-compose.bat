@echo=off

SET MONGODB_HOST=localhost
SET MONGODB_USERNAME=testuser
SET MONGODB_PASSWORD=yourpassword
SET PGSQL_PASSWORD=postgres
SET PGSQL_USER=postgres
SET PGSQL_DB=altastoredb

docker container rm -f mongodb-4.2
docker container rm -f postgresdb-10
docker container rm -f altastore-webservice2

docker image rm -f altastore-webservice2:1.0

docker build -t altastore-webservice2:1.0 .

docker-compose up -d
