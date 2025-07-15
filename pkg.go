// Package testgopkg
package testgopkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func HelloWorld() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	fmt.Printf("Hello, World\n")
}
