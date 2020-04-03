package main

import (
	"github.com/limetext/backend/log"
	_ "github.com/limetext/commands"
	"github.com/limetext/util"
	_ "github.com/zoli/sublime"
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

	err = fe.init()
	if err != nil {
		log.Error(err)
		return
	}

	defer shutDown()
	fe.loop()
}

func shutDown() {
	fe.shutDown()

	defer log.Close()
	log.Debug(util.Prof)
	if err := recover(); err != nil {
		log.Critical(err)
	}
}
