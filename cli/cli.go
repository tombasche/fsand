package cli

import (
	"log"
	"os"
)

type Args struct {
	Filename  string
	ToExecute string
}

func NewArgs() Args {
	if len(os.Args) != 3 {
		log.Fatalln("error")
	}
	filename := os.Args[1]
	toExecute := os.Args[2]
	log.Printf("Running '%s' on change", toExecute)
	return Args{Filename: filename, ToExecute: toExecute}
}
