package data

type Moon struct {
    Name              string  `json:"name"`
    ParentPlanet      string  `json:"parent_planet"`
    MassKg            float64 `json:"mass_kg"`
    DiameterKm        float64 `json:"diameter_km"`
    OrbitalPeriodDays float64 `json:"orbital_period_days"`
    Description       string  `json:"description"`
}