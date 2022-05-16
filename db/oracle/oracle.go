package oracle

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora"
)

// var goldDB = map[string]string{
// 	"service":  "CENTRAL_DEV",
// 	"username": "TOMGOLD_INT",
// 	"server":   "10.190.6.132",
// 	"port":     "1521",
// 	"password": "TOMGOLD_INT",
// }
type resulStr struct {
	IntId        string
	IntDate      string
	IntTimestamp string
	IntName      string
	IntData      string
}

func OracleConnection(dbParams map[string]string) *sql.DB {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"] //+ "?proxy client name=TOMGOLD_INT"

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}
	fmt.Println("Successfully connected!")
	return db
}
func MdmResult(db *sql.DB, interface_id string) resulStr {
	var int_data resulStr
	err := db.QueryRow("select * from j_interface_monitoring where interface_id = '"+interface_id+"'").
		Scan(&int_data.IntId, &int_data.IntDate, &int_data.IntTimestamp, &int_data.IntName, &int_data.IntData)
	if err != nil {
		fmt.Println(err)
	}

	return int_data

}
