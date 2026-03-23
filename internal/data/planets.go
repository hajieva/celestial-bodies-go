package data

type Atmosphere struct {
	Gas        string  `json:"gas"`
	Percentage float64 `json:"percentage"`
}

type Temperature struct {
	MinCelsius  float64 `json:"min_celsius"`
	MaxCelsius  float64 `json:"max_celsius"`
	MeanCelsius float64 `json:"mean_celsius"`
}

type Planet struct {
	Name                string       `json:"name"`
	Type                string       `json:"type"` // e.g. "Terrestrial", "Gas Giant", "Ice Giant"
	MassKg              float64      `json:"mass_kg"`
	DiameterKm          float64      `json:"diameter_km"`
	DistanceFromSunKm   float64      `json:"distance_from_sun_km"`
	OrbitalPeriodDays   float64      `json:"orbital_period_days"`
	RotationPeriodHours float64      `json:"rotation_period_hours"`
	Atmosphere          []Atmosphere `json:"atmosphere"`
	Temperature         Temperature  `json:"temperature"`
	MoonsCount          int          `json:"moons_count"`
	Description         string       `json:"description"`
}
