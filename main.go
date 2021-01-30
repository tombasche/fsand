package main

import (
	"fsand/alerting"
	"fsand/cli"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {

	args := cli.NewArgs()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	err = watcher.Add(args.Filename)
	if err != nil {
		log.Fatalln(err)
	}

	errCh := make(chan error)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					alerting.Alert(event.Op.String(), event.Name)
				}
			case err := <-watcher.Errors:
				errCh <- err
			}
		}
	}()

	log.Fatalln(<-errCh)
}
