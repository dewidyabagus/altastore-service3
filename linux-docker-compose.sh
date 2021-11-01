export MONGODB_HOST=localhost
export MONGODB_USERNAME=testuser
export MONGODB_PASSWORD=yourpassword
export PGSQL_PASSWORD=postgres
export PGSQL_USER=postgres
export PGSQL_DB=altastoredb

docker container rm -f mongodb-4.2
docker container rm -f postgresdb-10
docker container rm -f altastore-webservice2

docker-compose down

docker image rm -f altastore-webservice2:1.0

docker build -t altastore-webservice2:1.0 .

docker-compose up -d
