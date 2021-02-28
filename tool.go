package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const tool_endpoint = "api/printer/tool"

type ToolStruct struct {
Tool0 struct {
  Actual float32 `json:"actual"`
  Offset float32    `json:"offset"`
  Target float32 `json:"target"`
} `json:"tool0"`
}
//
func Status() ToolStruct {
	resTemp := Get(tool_endpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m ToolStruct
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

	}
	return m
}

func GetTemp(temp string) float32 {
	switch temp {
	case "offset":
		return Status().Tool0.Offset
	case "target":
		return Status().Tool0.Target
	default:
		return Status().Tool0.Actual
	}

}
