package infoparser

// This struct is used to store the global info about the API
type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}
