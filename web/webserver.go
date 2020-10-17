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

var HOST_NM, _ = os.Hostname()
var SERVER_NAME = os.Getenv("SERVER_NAME")

const SHORT_FORMAT = "2006-01-01 15:04:05"
const SERVER_PORT = ":8000"

var server *http.Server

func StartWebServer() {
	log.Println("please visit http://127.0.0.1" + SERVER_PORT)

	http.HandleFunc("/", root)
	http.HandleFunc("/time", showtime)
	http.HandleFunc("/sleep", sleep)

	server = &http.Server{
		Addr:           SERVER_PORT,
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
	log.Println("webs erver will stop")
	server.Shutdown(context.Background())
}

func root(writer http.ResponseWriter, request *http.Request) {
	logReqInfo(request)

	WriteFlush(writer, "response from ", HOST_NM, SERVER_NAME)

	WriteFlush(writer, "you can visit list below")
	WriteFlush(writer, "http://127.0.0.1:12345/time")
	WriteFlush(writer, "http://127.0.0.1:12345/sleep?cnt=10")

	showReqInfo(request,writer)
}

func showtime(writer http.ResponseWriter, request *http.Request) {
	logReqInfo(request)
	WriteFlush(writer, "response from ", HOST_NM, SERVER_NAME)
	s := fmt.Sprintf("hello,now time is %s", time.Now().Format(SHORT_FORMAT))
	WriteFlush(writer, s)
	log.Printf("%v\n", s)

	showReqInfo(request,writer)
}

func sleep(writer http.ResponseWriter, request *http.Request) {
	logReqInfo(request)
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

	showReqInfo(request,writer)
}

func WriteFlush(writer http.ResponseWriter, a ...interface{}) {
	fmt.Fprint(writer, a)
	fmt.Fprint(writer, "\n")
	writer.(http.Flusher).Flush()
}

func showReqInfo(request *http.Request,writer http.ResponseWriter){
	fmt.Fprint(writer, "\n")
	fmt.Fprint(writer,"------------req info ---------\n")
	fmt.Fprint(writer,"req url info:\n")
	fmt.Fprint(writer,fmt.Sprintf("    [%s]\n",request.URL.String()))
	fmt.Fprint(writer,"Header info:\n")
	for key, value := range request.Header {
		fmt.Fprint(writer,fmt.Sprintf("    [%s]=%v\n",key,value))
	}
	writer.(http.Flusher).Flush()
}

func logReqInfo(request *http.Request){
	log.Println("------------req info ---------")
	fmt.Println("Url info:")
	fmt.Printf("    [%s]\n",request.URL.String())
	fmt.Println("Header info:")
	for key, value := range request.Header {
		fmt.Printf("    [%s]=%v\n",key,value)
	}
}