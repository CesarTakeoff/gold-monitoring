package main

import (
	"fmt"

	"takeoff.com/monitoring/db/oracle"
	"takeoff.com/monitoring/db/postgres"
)

var goldDBConnection = map[string]string{
	"service":  "CENTRAL_DEV",
	"username": "TOMGOLD_INT",
	"server":   "10.190.6.132",
	"port":     "1521",
	"password": "TOMGOLD_INT",
}

var monitoringDBConnection = map[string]string{
	"host":     "104.198.38.44",
	"port":     "5432",
	"user":     "postgres",
	"password": `StYo{)X[";GR<**N`,
	"dbname":   "gold_monitoring_interfaces",
}

func main() {

	goldDB := oracle.OracleConnection(goldDBConnection)
	monitoringDB := postgres.PostgressConnection(monitoringDBConnection)
	defer monitoringDB.Close()
	defer goldDB.Close()

	//today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	//fmt.Println(today, today.String())

	intData := oracle.MdmResult(goldDB, "220414")
	fmt.Printf("Successful connection. Current data is: %v\n", intData)
	values := postgres.ResulStr{
		IntDate:      intData.IntDate,
		IntTimestamp: intData.IntTimestamp,
		IntId:        intData.IntId,
		IntName:      intData.IntName,
		IntData:      intData.IntData,
	}
	//fmt.Printf(values.IntData)
	postgres.AddMonitoringResul(monitoringDB, values)
	// rec := postgres.GetResult(monitoringDB)
	// fmt.Println(rec)
}
