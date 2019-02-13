// Binary towel runs a program specified on the command line, intercepts stderr, and reformats it to fatal structured
// log output.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must specify a program to run")
	}

	program := os.Args[1]
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:len(os.Args)]
	}

	c := exec.Command(program, args...)
	c.Stdout = os.Stdout
	stderrPipe, err := c.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Start(); err != nil {
		log.Fatal(err)
	}

	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Wait(); err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			now := time.Now().UTC()
			fmt.Printf(tmpl, now.Format(fmtTimePrefix), fmtTimeTS(now), program, program, err.Pid(), err.Error(), strconv.Quote(string(stderr)))
		default:
			log.Fatal(err)
		}
	}
}

// format: 2019-02-13T20:20:24.85271 {"level":"info","ts":1550089224.8527024,"logger":"foo","caller":"bar.go:1234","msg":"The log message","pid":11976,"json":{"arg1":"foo","arg2":bar}}
const tmpl = `%s {"level":"fatal","ts":%s,"logger":"%s","caller":"%s","msg":"Process ended with non-zero exit code","pid":%d,"json":{"error":%s,"stderr":%s}}`

const fmtTimePrefix = "2006-01-02T15:04:05.00000"

func fmtTimeTS(t time.Time) string {
	nanos := t.UnixNano()
	return fmt.Sprintf("%d.%d", nanos/10000000, nanos%10000000)
}
