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

	renameCh := make(chan bool)
	removeCh := make(chan bool)
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
	filename := args.Filename
	go func() {
		for {
			select {
			case <-renameCh:
				err = cli.WaitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			case <-removeCh:
				err = cli.WaitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()

	log.Fatalln(<-errCh)
}
