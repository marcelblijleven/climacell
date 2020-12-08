package climacell

import "testing"

func Test_field_String(t *testing.T) {
	tests := []struct {
		name string
		f    field
		want string
	}{
		{
			name: "Temperature",
			f:    Temperature,
			want: "temp",
		},
		{
			name: "FeelsLike",
			f:    FeelsLike,
			want: "feels_like",
		},
		{
			name: "DewPoint",
			f:    DewPoint,
			want: "dewpoint",
		},
		{
			name: "Humidity",
			f:    Humidity,
			want: "humidity",
		},
		{
			name: "WindSpeed",
			f:    WindSpeed,
			want: "wind_speed",
		},
		{
			name: "WindDirection",
			f:    WindDirection,
			want: "wind_direction",
		},
		{
			name: "WindGust",
			f:    WindGust,
			want: "wind_gust",
		},
		{
			name: "BarometricPressure",
			f:    BarometricPressure,
			want: "baro_pressure",
		},
		{
			name: "Precipitation",
			f:    Precipitation,
			want: "precipitation",
		},
		{
			name: "PrecipitationType",
			f:    PrecipitationType,
			want: "precipitation_type",
		},
		{
			name: "PrecipitationProbability",
			f:    PrecipitationProbability,
			want: "precipitation_probability",
		},
		{
			name: "PrecipitationAccumulation",
			f:    PrecipitationAccumulation,
			want: "precipitation_accumulation",
		},
		{
			name: "Sunrise",
			f:    Sunrise,
			want: "sunrise",
		},
		{
			name: "Sunset",
			f:    Sunset,
			want: "sunset",
		},
		{
			name: "Visibility",
			f:    Visibility,
			want: "visibility",
		},
		{
			name: "CloudCover",
			f:    CloudCover,
			want: "cloud_cover",
		},
		{
			name: "CloudBase",
			f:    CloudBase,
			want: "cloud_base",
		},
		{
			name: "CloudCeiling",
			f:    CloudCeiling,
			want: "cloud_ceiling",
		},
		{
			name: "CloudSatellite",
			f:    CloudSatellite,
			want: "cloud_satellite",
		},
		{
			name: "SurfaceShortwaveRadiation",
			f:    SurfaceShortwaveRadiation,
			want: "surface_shortwave_radiation",
		},
		{
			name: "MoonPhase",
			f:    MoonPhase,
			want: "moon_phase",
		},
		{
			name: "WeatherCode",
			f:    WeatherCode,
			want: "weather_code",
		},
		{
			name: "WeatherGroups",
			f:    WeatherGroups,
			want: "weather_groups",
		},
		{
			name: "ParticleMatter25",
			f:    ParticleMatter25,
			want: "pm25",
		},
		{
			name: "ParticleMatter10",
			f:    ParticleMatter10,
			want: "pm10",
		},
		{
			name: "Ozone",
			f:    Ozone,
			want: "o3",
		},
		{
			name: "NitrogenDioxide",
			f:    NitrogenDioxide,
			want: "no2",
		},
		{
			name: "CarbonMonoxide",
			f:    CarbonMonoxide,
			want: "co",
		},
		{
			name: "SulfurDioxide",
			f:    SulfurDioxide,
			want: "so2",
		},
		{
			name: "AirQualityIndexEPA",
			f:    AirQualityIndexEPA,
			want: "epa_aqi",
		},
		{
			name: "PrimaryPollutantEPA",
			f:    PrimaryPollutantEPA,
			want: "epa_primary_pollutant",
		},
		{
			name: "HealthConcernEPA",
			f:    HealthConcernEPA,
			want: "epa_health_concern",
		},
		{
			name: "AirQualityIndexChinaMEP",
			f:    AirQualityIndexChinaMEP,
			want: "china_aqi",
		},
		{
			name: "PrimaryPollutantChinaMEP",
			f:    PrimaryPollutantChinaMEP,
			want: "china_primary_pollutant",
		},
		{
			name: "HealthConcertChinaMEP",
			f:    HealthConcertChinaMEP,
			want: "china_health_concern",
		},
		{
			name: "TreePollen",
			f:    TreePollen,
			want: "pollen_tree",
		},
		{
			name: "WeedPollen",
			f:    WeedPollen,
			want: "pollen_weed",
		},
		{
			name: "GrassPollen",
			f:    GrassPollen,
			want: "pollen_grass",
		},
		{
			name: "RoadRiskScore",
			f:    RoadRiskScore,
			want: "road_risk_score",
		},
		{
			name: "RoadRisk",
			f:    RoadRisk,
			want: "road_risk",
		},
		{
			name: "RoadRiskConfidence",
			f:    RoadRiskConfidence,
			want: "road_risk_confidence",
		},
		{
			name: "RoadRiskConditions",
			f:    RoadRiskConditions,
			want: "road_risk_conditions",
		},
		{
			name: "FireIndex",
			f:    FireIndex,
			want: "fire_index",
		},
		{
			name: "HailBinary",
			f:    HailBinary,
			want: "hail_binary",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
