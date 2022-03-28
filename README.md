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
# Usage 
## Start server
Run the command that starts the server on host `127.0.0.1` and port `8080`.
```bash
go run main.go
```
After that open another terminal window and insure the path of current working directory is `*/api-go-gin-viper/src/`
___
## Add data block
Adds a data to the storage. Assuming that *address* and *amount* keys/values in `addBlock.json`
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
