package logging

import (
	"log"
)

func StdLog() (infoLog, errorLog *log.Logger) {
	// create info log
	fileInfo, err := openLogFile("./logs/myinfo.log")
	if err != nil {
		log.Fatal(err)
	}
	infoLog = log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Println("log info created")

	// create error log
	fileError, err := openLogFile("./logs/myerror.log")
	if err != nil {
		log.Fatal(err)
	}
	errorLog = log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Println("this is error")
	return infoLog, errorLog
}
