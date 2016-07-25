package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type data struct {
	//2016-05-25T21:22:00Z,120,0,4,9915,1,0
	AbsoluteTime      time.Time
	Number            int
	YAW               int
	Pitch             int
	VMC               int
	AmbientLightLevel int
	ActivityMask      int
}

func parseData(line string) data {
	arr := strings.Split(line, ",")
	ret := data{}
	t, _ := time.Parse(time.RFC3339, arr[0])
	ret.AbsoluteTime = t
	ret.Number = forceAtoi(arr[1])
	ret.YAW = forceAtoi(arr[2])
	ret.Pitch = forceAtoi(arr[3])
	ret.VMC = forceAtoi(arr[4])
	ret.AmbientLightLevel = forceAtoi(arr[5])
	ret.ActivityMask = forceAtoi(arr[6])
	return ret
}

func forceAtoi(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

type errorResponse struct {
	Message string
	Status  string
}

func newError(message string) []byte {
	tmp := errorResponse{
		Message: message,
		Status:  "error",
	}
	b, _ := json.Marshal(tmp)
	return b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "under construction")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(w, newError("Empty body."))
			return
		}
		d := parseData(string(b))
		ret, _ := json.Marshal(d)
		fmt.Fprint(w, string(ret))
	})
	port := getPort()
	log.Printf("listen on port %s", port)
	http.ListenAndServe(port, nil)
}

func getPort() string {
	env := os.Getenv("PORT")
	if env == "" {
		env = "8080" // for development
	}
	return ":" + env
}
