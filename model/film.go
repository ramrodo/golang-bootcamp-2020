package model

// Film - definition
// By adding struct tags you can control exactly what an how your struct will be marshalled to JSON
type Film struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
	ReleaseDate string `json:"releaseDate"`
	RtScore     string `json:"rtScore"`
}

// Films .
type Films []Film
