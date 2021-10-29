package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "log"
)

func main() {

    go func() {
        if r := recover(); r != nil {
            fmt.Printf(`err recovered, %s\r\n`, r)
        }
        return
    }()

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Panic(err)
    }
    defer watcher.Close()

    done := make(chan bool)

    go func() {
        for {
            select {
            case event := <-watcher.Events:
                fmt.Println(event)

            case err := <-watcher.Errors:
                fmt.Println("ERROR", err)
            }
        }
    }()

    if err := watcher.Add(`E:\Projects\project-github\daphne\practice`); err != nil {
        log.Panic(err)
    }
    <-done
}
