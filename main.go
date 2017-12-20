package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var port string
const version = "1"

// Probe acks back details about the container
func Probe(w http.ResponseWriter, req *http.Request) {
	h, _ := os.Hostname()
	s := "Hello from container: " + h + "\n"

	cmd := "ifconfig eth0 | grep 'inet addr' | cut -d: -f2 | awk '{ print $1}'"

	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Printf("Failed to execute command: %s", cmd)
	}
	s += "Local IP of container: " + string(out) + "\n"
	s += "Application version: " + version + "\n"
	io.WriteString(w, s)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", Probe)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	flag.StringVar(&port, "port", "9000", "Specify the port number the server will listen on.")
}
