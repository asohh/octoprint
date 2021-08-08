package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type JobStruct struct {
	Job struct {
		File struct {
			Name   string  `json:"name"`
			Origin string  `json:"origin"`
			Size   float32 `json:"size"`
			Date   float32 `json:"date"`
		} `json:"file"`
		EstimatedPrintTime float32 `json:"estimatedPrintTime"`
		Filament           struct {
			Length float32 `json:"length"`
			Volume float64 `json:"volume"`
		} `json:"filament"`
	} `json:"job"`
	Progress struct {
		Completion    float32 `json:"completion"`
		Filepos       float32 `json:"filepos"`
		PrintTime     float32 `json:"printTime"`
		PrintTimeLeft float32 `json:"printTimeLeft"`
	} `json:"progress"`
	State string `json:"state"`
}


const jobEndpoint = "/api/job"

func execute(command string, action string) {
	switch command {
	case "pause":
		Post(jobEndpoint, []byte("{\"command\":\""+command+"\",\"action\": \""+action+"\"}"))
	default:
		Post(jobEndpoint, []byte("{\"command\":\""+command+"\"}"))
	}
}

func JobCancel() {
	execute("cancel", "")
}

func JobStart() {
	execute("start", "")
}

func JobRestart() {
	execute("restart", "")
}

func JobPause(action string) {
	switch action {
	case "pause", "resume":
		execute("pause", action)
	default:
		execute("pause", "toggle")
	}

}
func JobTimeRemaining() float32{
return JobStatus().Progress.PrintTimeLeft

}
func JobStatus() JobStruct {
	resTemp := Get(jobEndpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m JobStruct
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

	}
	return m

}
