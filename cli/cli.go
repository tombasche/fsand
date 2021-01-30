package cli

import (
	"log"
	"os"
	"time"
)

func WaitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}

type Args struct {
	Filename  string
	ToExecute string
}

func NewArgs() Args {
	if len(os.Args) != 3 {
		log.Fatalln("error")
	}
	filename := os.Args[1]
	err := WaitUntilFind(filename)
	if err != nil {
		log.Fatalln(err)
	}
	toExecute := os.Args[2]
	log.Printf("Running '%s' on change", toExecute)
	return Args{Filename: filename, ToExecute: toExecute}
}
