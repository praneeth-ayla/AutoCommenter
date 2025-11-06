package scanner

type FileInfo struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Lines int    `json:"lines"`
	Size  int64  `json:"size"`
}
