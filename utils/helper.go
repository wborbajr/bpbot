package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var GeneralLogger *log.Logger

var ErrorLogger *log.Logger

func init() {

	logDir := "log"

	_ , err := os.Stat(logDir)

	if os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)

		if err != nil {
			fmt.Println("Error creating log folder")
		}
	}

	absPath, err := filepath.Abs(logDir)
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/general-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
