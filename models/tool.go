package octoprint

type ToolStruct struct {
Tool0 struct {
  Actual float32 `json:"actual"`
  Offset float32    `json:"offset"`
  Target float32 `json:"target"`
} `json:"tool0"`
}
