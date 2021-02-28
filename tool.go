package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const toolEndpoint = "api/printer/tool"

type ToolStruct struct {
Tool0 struct {
  Actual float32 `json:"actual"`
  Offset float32    `json:"offset"`
  Target float32 `json:"target"`
} `json:"tool0"`
}
//
func ToolStatus() ToolStruct {
	resTemp := Get(toolEndpoint)
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

func ToolGetTemp(temp string) float32 {
	switch temp {
	case "offset":
		return ToolStatus().Tool0.Offset
	case "target":
		return ToolStatus().Tool0.Target
	default:
		return ToolStatus().Tool0.Actual
	}

}
