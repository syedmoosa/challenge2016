package model

// Permission - Contains permissions given to the distributor
type Permission struct {
	For      string     `json:"for,omitempty"`
	Includes []string   `json:"includes,omitempty"`
	Excludes []Excluded `json:"excludes,omitempty"`
}

//Excluded - Contains excluded city, province and country for the distributor
type Excluded struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Country  string `json:"country,omitempty"`
}
