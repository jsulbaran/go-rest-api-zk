package service

import (
	"database/sql"
	"log"
	"os"
)

func GetDeviceSerial(systemDbPath string) string {
	if _, err := os.Stat(systemDbPath); os.IsNotExist(err) {
		return "error"
	}

	database, err := sql.Open("sqlite3", systemDbPath)
	if err != nil {
		log.Fatalf("Error de bd sistema: %v", err)
		panic(err.Error())
		return ""
	}
	defer database.Close()
	var serialNumber string
	row := database.QueryRow("select Factory_Set_Value as serial_number from OPTION_INFO where Option_Name ='~SerialNumberEx'")
	switch err := row.Scan(&serialNumber); err {
	case sql.ErrNoRows:
		//fmt.Println("No rows were returned!")
	case nil:
		//fmt.Println(id)
	default:
		return "error"
	}
	return serialNumber

}
