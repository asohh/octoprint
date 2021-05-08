package octoprint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type FilesStruct struct {
	Files []struct {
		Name     string   `json:"name"`
		Path     string   `json:"path"`
		Type     string   `json:"type"`
		TypePath []string `json:"typePath"`
		Hash     string   `json:"hash,omitempty"`
		Size     int      `json:"size,omitempty"`
		Date     int      `json:"date,omitempty"`
		Origin   string   `json:"origin,omitempty"`
		Refs     struct {
			Resource string `json:"resource"`
			Download string `json:"download"`
		} `json:"refs,omitempty"`
		GcodeAnalysis struct {
			EstimatedPrintTime int `json:"estimatedPrintTime"`
			Filament           struct {
				Length int     `json:"length"`
				Volume float64 `json:"volume"`
			} `json:"filament"`
		} `json:"gcodeAnalysis,omitempty"`
		Print struct {
			Failure int `json:"failure"`
			Success int `json:"success"`
			Last    struct {
				Date    int  `json:"date"`
				Success bool `json:"success"`
			} `json:"last"`
		} `json:"print,omitempty"`
		Children []struct {
			Name     string   `json:"name"`
			Path     string   `json:"path"`
			Type     string   `json:"type"`
			TypePath []string `json:"typePath"`
			Hash     string   `json:"hash"`
			Size     int      `json:"size"`
			Date     int      `json:"date"`
			Origin   string   `json:"origin"`
			Refs     struct {
				Resource string `json:"resource"`
				Download string `json:"download"`
			} `json:"refs"`
			GcodeAnalysis struct {
				EstimatedPrintTime int `json:"estimatedPrintTime"`
				Filament           struct {
					Length int     `json:"length"`
					Volume float64 `json:"volume"`
				} `json:"filament"`
			} `json:"gcodeAnalysis"`
			Print struct {
				Failure int `json:"failure"`
				Success int `json:"success"`
				Last    struct {
					Date    int  `json:"date"`
					Success bool `json:"success"`
				} `json:"last"`
			} `json:"print"`
		} `json:"children,omitempty"`
	} `json:"files"`
	Free string `json:"free"`
}


const filesEndpoint = "/api/FilesStruct"

func SelectFile(path string, print string) {
		Post(filesEndpoint, []byte("{\"command\":\"select\",\"print\":"+print+"}"))
}

func FilesStatus() FilesStruct {
	resTemp := Get(filesEndpoint)
	jsonTemp, _ := ioutil.ReadAll(resTemp.Body)
	resTemp.Body.Close()
	dec := json.NewDecoder(strings.NewReader(string(jsonTemp)))
	var m FilesStruct
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

	}
	return m

}
