package octoprint

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
