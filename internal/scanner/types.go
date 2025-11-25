package scanner

// Info represents metadata for a scanned file.
type Info struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Lines int    `json:"lines"`
	Size  int64  `json:"size"`
}

type Data struct {
	Path    string
	Content string
}