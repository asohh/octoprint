package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const endpoint = "/api/job"

func execute(command string, action string) {
	switch command {
	case "pause":
		api.Post(endpoint, []byte("{\"command\":\""+command+"\",\"action\": \""+action+"\"}"))
	default:
		api.Post(endpoint, []byte("{\"command\":\""+command+"\"}"))
	}
}

func Cancel() {
	execute("cancel", "")
}

func Start() {
	execute("start", "")
}

func Restart() {
	execute("restart", "")
}

func Pause(action string) {
	switch action {
	case "pause", "resume":
		execute("pause", action)
	default:
		execute("pause", "toggle")
	}

}

func Status() models.JobStruct {
	resTemp := api.Get(endpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m models.JobStruct
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

	}
	return m

}
