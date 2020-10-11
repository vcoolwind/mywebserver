package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var server *http.Server

var HOST_NM, _ = os.Hostname()
var SERVER_NAME = os.Getenv("SERVER_NAME")

func StartWebServer() {
	log.Println("please visit http://127.0.0.1:12345")

	http.HandleFunc("/", root)
	http.HandleFunc("/time", showtime)
	http.HandleFunc("/sleep", sleep)

	server = &http.Server{
		Addr:           ":12345",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err == nil {
		log.Println("http server running")
	} else {
		log.Println("http server run fail:", err)
		os.Exit(-1)
	}
}

func Stop() {
	log.Println("webserver will stop")
	server.Shutdown(context.Background())
}

func root(writer http.ResponseWriter, request *http.Request) {
	WriteFlush(writer, "response from ", HOST_NM, SERVER_NAME)

	WriteFlush(writer, "you can visit list below")
	WriteFlush(writer, "http://127.0.0.1:12345/time")
	WriteFlush(writer, "http://127.0.0.1:12345/sleep?cnt=10")
}

func showtime(writer http.ResponseWriter, request *http.Request) {
	WriteFlush(writer, "response from ", HOST_NM, SERVER_NAME)
	const shortForm = "2006-01-01 15:04:05"
	s := fmt.Sprintf("hello ,now time is %s", time.Now().Format(shortForm))
	fmt.Fprint(writer, s)
	log.Printf("%v\n", s)
}

func sleep(writer http.ResponseWriter, request *http.Request) {
	WriteFlush(writer, "response from ", HOST_NM, SERVER_NAME)
	var info, cnt string
	cnt = request.URL.Query().Get("cnt")
	if cnt == "" {
		cnt = "3"
	}
	sleepCnt, _ := strconv.Atoi(cnt)
	info = "will sleep " + strconv.Itoa(sleepCnt)
	log.Println(info)
	WriteFlush(writer, info)
	for i := 0; i < sleepCnt; i++ {
		time.Sleep(1 * time.Second)
		WriteFlush(writer, fmt.Sprintf("sleep %d second", i+1))
	}
	info = "sleep over"
	WriteFlush(writer, info)
	log.Printf("%v\n", info)
}

func WriteFlush(writer http.ResponseWriter, a ...interface{}) {
	fmt.Fprint(writer, a)
	fmt.Fprint(writer, "\n")
	writer.(http.Flusher).Flush()
}
