package climacell

type field int

const (
	// Core
	Temperature field = iota
	FeelsLike
	DewPoint
	Humidity
	WindSpeed
	WindDirection
	WindGust
	BarometricPressure
	Precipitation
	PrecipitationType
	PrecipitationProbability
	PrecipitationAccumulation
	Sunrise
	Sunset
	Visibility
	CloudCover
	CloudBase
	CloudCeiling
	CloudSatellite
	SurfaceShortwaveRadiation
	MoonPhase
	WeatherCode
	WeatherGroups

	// Air quality
	ParticleMatter25
	ParticleMatter10
	Ozone
	NitrogenDioxide
	CarbonMonoxide
	SulfurDioxide
	AirQualityIndexEPA
	PrimaryPollutantEPA
	HealthConcernEPA
	AirQualityIndexChinaMEP
	PrimaryPollutantChinaMEP
	HealthConcertChinaMEP

	// Pollen
	TreePollen
	WeedPollen
	GrassPollen

	// Road
	RoadRiskScore
	RoadRisk
	RoadRiskConfidence
	RoadRiskConditions

	// Fire
	FireIndex

	// Insurance
	HailBinary
)

var fieldValues = []string{
	"temp",
	"feels_like",
	"dewpoint",
	"humidity",
	"wind_speed",
	"wind_direction",
	"wind_gust",
	"baro_pressure",
	"precipitation",
	"precipitation_type",
	"precipitation_probability",
	"precipitation_accumulation",
	"sunrise",
	"sunset",
	"visibility",
	"cloud_cover",
	"cloud_base",
	"cloud_ceiling",
	"cloud_satellite",
	"surface_shortwave_radiation",
	"moon_phase",
	"weather_code",
	"weather_groups",
	"pm25",
	"pm10",
	"o3",
	"no2",
	"co",
	"so2",
	"epa_aqi",
	"epa_primary_pollutant",
	"epa_health_concern",
	"china_aqi",
	"china_primary_pollutant",
	"china_health_concern",
	"pollen_tree",
	"pollen_weed",
	"pollen_grass",
	"road_risk_score",
	"road_risk",
	"road_risk_confidence",
	"road_risk_conditions",
	"fire_index",
	"hail_binary",
}

// String returns the string value of the unit
func (f field) String() string {
	return fieldValues[f]
}
