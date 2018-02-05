package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func pullHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		shFile := "./secret-sh-file.sh"
		_, stderr := shellExec("/bin/bash", shFile)
		res := &response{200, stderr}
		res.json(w)
	}
	return
}

func shellExec(args ...string) (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	cmd := exec.Command(args[0], args[1:]...)

	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	errStr, _ := ioutil.ReadAll(stderr)
	outStr, _ := ioutil.ReadAll(stdout)

	return string(outStr), string(errStr)
}
