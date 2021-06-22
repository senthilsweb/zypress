package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

var stage string

func init() {
	log.Info("Application init function start")
	initLogger()
	log.Info("Application init function end.")
}

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {

	port := flag.String("p", "8100", "port to serve on")
	indexDoc := flag.String("i", "index.html", "index document to load")
	directory := flag.String("d", "dist", "the directory of static file to web host")
	env := flag.String("e", "dev", "dev or prod enviroment?")
	stage = *env
	log.Info(env, port, indexDoc, directory)
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: *directory, indexPath: *indexDoc}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Server starting at port " + *port)

	//log.Fatal(http.ListenAndServe(":"+port, a.Negroni))
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go http.Serve(listener, router)
	fmt.Println("Started Server at port " + *port)
	<-done
}

func initLogger() {
	logfilepath := AppExecutionPath() + "/" + os.Args[0] + ".log"
	log.Info("logfilepath = " + logfilepath)
	// Set the Lumberjack logger
	ljack := &lumberjack.Logger{
		Filename:   logfilepath,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     3,
		LocalTime:  true,
	}

	//log := logrus.New()
	//
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	mWriter := io.MultiWriter(os.Stdout, ljack)
	log.SetOutput(mWriter)
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"Runtime Version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application Initializing")

	if stage == "Dev" {
		log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.RFC1123Z})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

// AppExecutionPath returns the relative path where the application is executing
func AppExecutionPath() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}
