package main

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

func init() {
	http.HandleFunc("/start", startTask)
	http.HandleFunc("/back/task", receiveTask)
}

func startTask(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	task := taskqueue.NewPOSTTask("/back/task", nil)
	for i := 0; i < 10000; i++ {
		taskqueue.Add(ctx, task, "queue1")
		taskqueue.Add(ctx, task, "queue2")
		taskqueue.Add(ctx, task, "queue3")
		taskqueue.Add(ctx, task, "queue4")
		taskqueue.Add(ctx, task, "queue5")
	}
}

func receiveTask(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "receive task")
	time.Sleep(50 * time.Millisecond)
}
