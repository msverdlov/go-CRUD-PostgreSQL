# Simple API implementation in Go

## To Get Started
#### Clone repo
1. Make sure you have [Go](https://go.dev/) installed on your computer.
2. Run the command to clone this github repository and change the directory to the project's folder:
```bash
git clone https://github.com/vsevdrob/api-go-gin-viper.git && cd api-go-gin-viper/src
```
#### Before we run the programme, let's download a couple of Go dependencies first.
3. This will download one of the coolest HTTP web framework [gin](https://github.com/gin-gonic/gin) written in Go (Golang).
```bash
go get github.com/gin-gonic/gin
```
4. [Viper](https://github.com/spf13/viper) helps us to to operate with the predefined `config.yaml` file in order to export some required values from it.
```bash
go get github.com/spf13/viper
```

5. Add to git
`git remote set-url origin git@github.com:msverdlov/go-CRUD-PostgreSQL.git`

# Usage
## Start server in 1-st bash
Run the command that starts the server on host `127.0.0.1` and port `8080`
```bash
cd cmd/anyData
go run main.go
```

## Start server in 2-nd bash
Rus PostgreSQL
``
docker-compose up --build
``

##In manual mode exec:
```
create table content
(
id          serial
constraint content_pk
primary key,
mata_data   varchar,
date_create timestamp,
date_update timestamp
);

alter table content
owner to "goUser";

create unique index content_id_uindex
on content (id);
```

___
## Add data block in 3-th bash
```bash
curl localhost:8080/add-data --header "Content-Type: application json" -d @examples/addBlock.json --request "POST"
```
## Fetch data
Get a data by assigning his/her **`id`** in the query.
```bash
curl localhost:8080/get-data?id=1 --request "GET"
```
## Fetch dataset
Get a add all content from the DB.
```bash
curl localhost:8080/get-content --request "GET"
```
## Update date
Update data amount by increasing it. Assign the **`id`** in the query to update a specific data.
```bash
curl localhost:8080/update-data --header "Content-Type: application json" -d @examples/updateBlock.json --request "PATCH"
```
## Delete data
Delete a data from DB. Assign data **`id`** in the query to delete the preferred one.
```bash
curl localhost:8080/delete-data?id=10 --request "DELETE"
```
# Licence
**MIT**
