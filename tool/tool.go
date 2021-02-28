package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const endpoint = "api/printer/tool"

//
func Status() models.ToolStruct {
	resTemp := api.Get(endpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m models.ToolStruct
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
