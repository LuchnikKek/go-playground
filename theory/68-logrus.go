// logrus - структурированное логирование (json)

package theory

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func logrusExample() {
	log.WithFields(log.Fields{
		"genre": "romance",
		"name":  "Rammstein",
	}).Info("Немецкая группа, основанная в 1994.")

	log.WithFields(log.Fields{
		"omg":  true,
		"name": "Garbage",
	}).Warn("В 2021 году вышел новый альбом No Gods No Masters.")
	//{"level":"warning","msg":"В 2021 году вышел новый альбом No Gods No Masters.","name":"Garbage","omg":true,"time":"2025-06-22T19:09:50+03:00"}

	log.WithFields(log.Fields{
		"omg":  true,
		"name": "Linkin Park",
	}).Error("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
	//{"level":"error","msg":"Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.","name":"Linkin Park","omg":true,"time":"2025-06-22T19:09:50+03:00"}
}

func logrusContextExample() {
	contextLogger := log.WithFields(log.Fields{
		"common": "Any music is awesome",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I will be logged with common and other fields")
	//{"common":"Any music is awesome","level":"warning","msg":"I will be logged with common and other fields","other":"I also should be logged always","time":"2025-06-22T19:16:02+03:00"}

	contextLogger.Error("Me too, maybe")
	//{"common":"Any music is awesome","level":"error","msg":"Me too, maybe","other":"I also should be logged always","time":"2025-06-22T19:16:02+03:00"}
}

func MainLogrus() {
	var buf bytes.Buffer

	log.SetOutput(&buf)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)

	logrusExample()
	logrusContextExample()

	fmt.Println(&buf)
}
