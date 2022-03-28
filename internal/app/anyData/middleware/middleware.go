package middleware

import (
	"anyData/internal/app/anyData/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "goUser"
	password = "!JustPwd4goUsr#74Anna"
	dbname   = "anyDataDB"
)

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("[MID-debug] Successfully connected to DB")

	return db
}

func dbClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		//log.Fatalf("Unable to close db connection: %v", err)
	}
}

func rowsClose(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		//log.Fatalf("Unable to close fetch rows: %v", err)
	}
}

func fetchRows(rows *sql.Rows) []model.AnyDataStruct  {
	var content []model.AnyDataStruct
	for rows.Next() {
		var row model.AnyDataStruct
		err := rows.Scan(&row.Id, &row.MetaData, &row.DateUpdate, &row.DateCreate)
		if err != nil {
			//log.Fatalf("Unable to scan the row. %v", err)
		}
		content = append(content, row)
	}

	if len(content) == 0 {
		return nil
	}

	return content
}

func GetContent() ([]model.AnyDataStruct, error) {
	db := createConnection()
	defer dbClose(db)

	sqlStatement := `SELECT * FROM content`
	rows, err := db.Query(sqlStatement)
	defer rowsClose(rows)

	return fetchRows(rows), err
}

func GetDataById(id uint64) (*model.AnyDataStruct, error) {
	db := createConnection()
	defer dbClose(db)

	sqlStatement := `SELECT * FROM content WHERE id=$1 LIMIT 1`
	rows, err := db.Query(sqlStatement, id)
	defer rowsClose(rows)

	row := fetchRows(rows)
	if row == nil {
		return nil, err
	}

	return &row[0], err
}

func AddData(data model.AnyDataStruct) uint64 {
	db := createConnection()
	defer dbClose(db)

	sqlStatement := `INSERT INTO content (mata_data, date_create, date_update) VALUES ($1, $2, $3) RETURNING id`
	var id uint64

	err := db.QueryRow(sqlStatement, data.MetaData, data.DateCreate, data.DateUpdate).Scan(&id)
	if err != nil {
		//log.Fatalf("Unable to execute the query. %v", err)
	}

	return id
}

func UpdateData(data model.AnyDataStruct) bool {
	db := createConnection()
	defer dbClose(db)

	sqlStatement := `UPDATE content SET mata_data=$2, date_update=$3 WHERE id=$1`
	res, err := db.Exec(sqlStatement, data.Id, data.MetaData, data.DateUpdate)
	if err != nil {
		//log.Fatalf("Unable to execute the query. %v", err)
		return false
	}
	rows, err := res.RowsAffected()
	if err != nil {
		//log.Fatalf("Error while checking the affected rows. %v", err)
		return false
	}
	if rows == 0 {
		//log.Fatalf("Error while checking the affected rows. %v", err)
		return false
	}

	return true
}

func DeleteData(id uint64) bool {
	db := createConnection()
	defer dbClose(db)
	sqlStatement := `DELETE FROM content WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		//log.Fatalf("Unable to execute the query. %v", err)
		return false
	}
	rows, err := res.RowsAffected()
	if err != nil {
		//log.Fatalf("Error while checking the affected rows. 111 %v", err)
		return false
	}
	if rows == 0 {
		//log.Fatalf("Error while checking the affected rows. 222 %v", err)
		return false
	}

	return true
}