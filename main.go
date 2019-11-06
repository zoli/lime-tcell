package main

import (
	"github.com/limetext/backend/log"
	"github.com/limetext/util"
)

var fe *frontend

func main() {
	var err error
	log.AddFilter("stdout", log.FINEST, log.NewFileLogWriter("debug.log", false))

	fe, err = newFrontend()
	if err != nil {
		log.Error(err)
		return
	}
	defer shutDown()

	fe.init()
	fe.editor.ActiveWindow().OpenFile("./main.go", 0)

	fe.loop()
}

func shutDown() {
	fe.shutDown()

	defer log.Close()
	log.Debug(util.Prof)
	if err := recover(); err != nil {
		log.Critical(err)
		panic(err)
	}
}
