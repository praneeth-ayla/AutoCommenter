package contextstore

type FileDetails struct {
	Path     string   `json:"path"`
	Name     string   `json:"file_name"`
	Exports  []string `json:"exports"`
	ImpLogic []string `json:"imp_logic"`
	Summary  string   `json:"summary"`
}
