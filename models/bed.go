package octoprint

type BedStruct struct {
	Bed struct {
		Actual float32 `json:"actual"`
		Offset float32 `json:"offset"`
		Target float32 `json:"target"`
	} `json:"bed"`
}
