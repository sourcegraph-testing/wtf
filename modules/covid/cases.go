package covid

// Cases holds the latest cases
type Cases struct {
	Latest Latest `json:"latest"`
}

// Latest holds the number of global confirmed cases and deaths due to Covid
type Latest struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	// Not currently used but holds information about the country
	Locations []any `json:"locations,omitempty"`
}
