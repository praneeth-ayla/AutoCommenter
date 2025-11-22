package contextstore

type FileDetails struct {
	Path    string   `json:"path"`
	Name    string   `json:"file_name"`
	Exports []string `json:"exports"`
	Imports []string `json:"imports"`
	Summary string   `json:"summary"`
}
