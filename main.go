package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	addr     string
	port     string
	hostname string
	s        string
	out      string
)

// Probe acks back details about the container
func Probe(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, s)
}

func main() {
	flag.Parse()
	hostname, _ := os.Hostname()
	s := "Hello from container: " + hostname + "\n"

	cmd := "ifconfig eth0 | grep 'inet addr' | cut -d: -f2 | awk '{ print $1}'"
	out, err := exec.Command("sh", "-c", cmd).Output()
	s += "Local IP of container: " + string(out) + "\n"

	if err != nil {
		log.Printf("Failed to execute command: %s", cmd)
	}

	http.HandleFunc("/", Probe)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	flag.StringVar(&port, "port", "9000", "Specify the port number the server will listen on.")
	flag.StringVar(&addr, "addr", "127.0.0.1", "Specify the address the server will listen on.")

}
