package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

var (
	BUF_LEN = 1024
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("restic", "-q", "--json", "cat", "blob", "a9083042452d72a1e36cdc3116c48d1e3330b572b0c7a6382b926a52f5dc1b59")
	// cmd := exec.Command("restic", "-q", "--json", "snapshots")
	// cmd := exec.Command("restic", "-q", "--json", "ls", "01380af8")

	env := os.Environ()
	env = append(env, "RESTIC_PASSWORD=test")
	env = append(env, "RESTIC_REPOSITORY=/tmp/backup")
	cmd.Env = env

	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter
	go writeCmdOutput(w, pipeReader)
	cmd.Run()
	pipeWriter.Close()
}

func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		res.Write(data)
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
