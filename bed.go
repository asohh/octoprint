package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const bed_endpoint = "api/printer/bed"

type BedStruct struct {
	Bed struct {
		Actual float32 `json:"actual"`
		Offset float32 `json:"offset"`
		Target float32 `json:"target"`
	} `json:"bed"`
}

//
func Status() BedStruct {
	resTemp := Get(bed_endpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m BedStruct
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
		return Status().Bed.Offset
	case "target":
		return Status().Bed.Target
	default:
		return Status().Bed.Actual
	}

}
