package contextstore

type FileDetails struct {
	Path    string   `json:"path"`      // Path to the file.
	Name    string   `json:"file_name"` // Name of the file.
	Exports []string `json:"exports"`   // List of exported identifiers from the file.
	Imports []string `json:"imports"`   // List of imported packages.
	Summary string   `json:"summary"`   // A brief summary of the file's purpose.
}
