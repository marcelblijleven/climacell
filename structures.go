package climacell

import (
	"encoding/json"
	"fmt"
	"time"
)

// RealtimeData embeds the CoreLayer, AirQualityLayer, PollenLayer, RoadLayer and FireLayer. It's the response type
// for the Realtime() API call
type RealtimeData struct {
	CoreLayer
	AirQualityLayer
	PollenLayer
	RoadLayer
	FireLayer
}

// NowcastData embeds the CoreLayer, AirQualityLayer, PollenLayer, RoadLayer and InsuranceLayer. It's the response type
// for the Nowcast() API call
type NowcastData struct {
	ApiResponse
	CoreLayer
	AirQualityLayer
	PollenLayer
	RoadLayer
	InsuranceLayer
}

// ApiResponse is the basic api response which contains only latitude, longitude and observation time
type ApiResponse struct {
	Latitude        float64  `json:"lat"`
	Longitude       float64  `json:"lon"`
	ObservationTime TimeData `json:"observation_time"`
}

type NowcastResponse struct {
	Timesteps []NowcastData
}

// CoreLayer is the data layer of type Core, which is used in
// all ClimaCell products (not applicable to each data field in the layer)
type CoreLayer struct {
	Temperature               *FloatData  `json:"temp,omitempty"`
	FeelsLike                 *FloatData  `json:"feels_like,omitempty"`
	DewPoint                  *FloatData  `json:"dew_point,omitempty"`
	WindSpeed                 *FloatData  `json:"wind_speed,omitempty"`
	WindGust                  *FloatData  `json:"wind_gust,omitempty"`
	BarometricPressure        *FloatData  `json:"baro_pressure,omitempty"`
	Visibility                *IntData    `json:"visibility,omitempty"`
	Humidity                  *FloatData  `json:"humidity,omitempty"`
	WindDirection             *FloatData  `json:"wind_direction,omitempty"`
	Precipitation             *FloatData  `json:"precipitation,omitempty"`
	PrecipitationType         *StringData `json:"precipitation_type,omitempty"`
	PrecipitationProbability  *FloatData  `json:"precipitation_probability,omitempty"`
	PrecipitationAccumulation *FloatData  `json:"precipitation_accumulation,omitempty"`
	CloudCover                *FloatData  `json:"cloud_cover,omitempty"`
	CloudCeiling              *IntData    `json:"cloud_ceiling,omitempty"`
	CloudBase                 *IntData    `json:"cloud_base,omitempty"`
	CloudSatellite            *FloatData  `json:"cloud_satellite,omitempty"`
	SurfaceShortwaveRadiation *IntData    `json:"surface_shortwave_radiation,omitempty"`
	Sunrise                   *TimeData   `json:"sunrise,omitempty"`
	Sunset                    *TimeData   `json:"sunset,omitempty"`
	MoonPhase                 *StringData `json:"moon_phase,omitempty"`
	WeatherCode               *StringData `json:"weather_code,omitempty"`
	WeatherGroups             *[]string   `json:"weather_groups,omitempty"`
}

// PollenLayer is the data layer of type Air quality, which is used in
// Realtime, Nowcast, Hourly, ClimaCell and Tiles
type AirQualityLayer struct {
	ParticulateMatter25      *FloatData  `json:"pm25,omitempty"`
	ParticulateMatter10      *FloatData  `json:"pm10,omitempty"`
	Ozone                    *FloatData  `json:"03,omitempty"`
	NitrogenDioxide          *FloatData  `json:"no2,omitempty"`
	CarbonMonoxide           *FloatData  `json:"co,omitempty"`
	SulfurDioxide            *FloatData  `json:"so2,omitempty"`
	AirQualityIndexEPA       *FloatData  `json:"epa_aqi,omitempty"`
	PrimaryPollutantEPA      *StringData `json:"epa_primary_pollutant,omitempty"`
	HealthConcernEPA         *StringData `json:"epa_health_concern,omitempty"`
	AirQualityIndexChinaMEP  *FloatData  `json:"china_aqi,omitempty"`
	PrimaryPollutantChinaMEP *StringData `json:"china_primary_pollutant,omitempty"`
	HealthConcernChinaMEP    *StringData `json:"china_health_concern,omitempty"`
}

// PollenLayer is the data layer of type Pollen, which is used in
// Realtime, Nowcast, Hourly, ClimaCell and Tiles
type PollenLayer struct {
	PollenTree  PollenTree  `json:"pollen_tree,omitempty"`
	PollenWeed  PollenWeed  `json:"pollen_weed,omitempty"`
	PollenGrass PollenGrass `json:"pollen_grass,omitempty"`
}

// RoadLayer is the data layer of type Road, which is used in
// Realtime, Nowcast, Hourly and ClimaCell
type RoadLayer struct {
	RoadRiskScore      string `json:"road_risk_score,omitempty"`
	RoadRisk           string `json:"road_risk,omitempty"`
	RoadRiskConfidence int    `json:"road_risk_confidence,omitempty"`
	RoadRiskConditions string `json:"road_risk_conditions,omitempty"`
}

// FireLayer is the data layer of type Fire, which is used in
// Realtime, ClimaCell and Tiles
type FireLayer struct {
	FireIndex *FloatData `json:"fire_index,omitempty"`
}

// InsuranceLayer is the data layer of type Insurance, which is used in
// Realtime, Nowcast, Hourly and ClimaCell
type InsuranceLayer struct {
	HailBinary *IntData `json:"hail_binary,omitempty"`
}

// PollenTree are the various trees that emit pollen when they're in season
type PollenTree struct {
	Acacia     *PollenData `json:"pollen_tree_acacia,omitempty"`
	Ash        *PollenData `json:"pollen_tree_ash,omitempty"`
	Beech      *PollenData `json:"pollen_tree_beech,omitempty"`
	Birch      *PollenData `json:"pollen_tree_birch,omitempty"`
	Cedar      *PollenData `json:"pollen_tree_cedar,omitempty"`
	Cypress    *PollenData `json:"pollen_tree_cypress,omitempty"`
	Elder      *PollenData `json:"pollen_tree_elder,omitempty"`
	Elm        *PollenData `json:"pollen_tree_elm,omitempty"`
	Hemlock    *PollenData `json:"pollen_tree_hemlock,omitempty"`
	Hickory    *PollenData `json:"pollen_tree_hickory,omitempty"`
	Juniper    *PollenData `json:"pollen_tree_juniper,omitempty"`
	Mahogany   *PollenData `json:"pollen_tree_mahogany,omitempty"`
	Maple      *PollenData `json:"pollen_tree_maple,omitempty"`
	Mulberry   *PollenData `json:"pollen_tree_mulberry,omitempty"`
	Oak        *PollenData `json:"pollen_tree_oak,omitempty"`
	Pine       *PollenData `json:"pollen_tree_pine,omitempty"`
	Cottonwood *PollenData `json:"pollen_tree_cottonwood,omitempty"`
	Spruce     *PollenData `json:"pollen_tree_spruce,omitempty"`
	Sycamore   *PollenData `json:"pollen_tree_sycamore,omitempty"`
	Walnut     *PollenData `json:"pollen_tree_walnut,omitempty"`
	Willow     *PollenData `json:"pollen_tree_willow,omitempty"`
}

// PollenWeed are the various weeds that emit pollen
type PollenWeed struct {
	Ragweed *PollenData `json:"pollen_weed_ragweed,omitempty"`
}

// PollenGrass are the various weeds that emit pollen
type PollenGrass struct {
	Grass *PollenData `json:"pollen_grass_grass,omitempty"`
}

// PollenData represents the pollen value
type PollenData struct {
	Value int `json:"value"`
}

// FloatData is a response type in which a float and unit are stored
type FloatData struct {
	Value *float64 `json:"value"`
	Units string   `json:"units,omitempty"`
}

// String represent the string value of FloatData
func (d *FloatData) String() string {
	if d.Value == nil {
		return fmt.Sprintf("0 %v", d.Units)
	}

	return fmt.Sprintf("%g %v", *d.Value, d.Units)
}

// IntData is a response type in which an int and unit are stored
type IntData struct {
	Value *int   `json:"value"`
	Units string `json:"units,omitempty"`
}

// String represent the string value of IntData
func (d *IntData) String() string {
	if d.Value == nil {
		return fmt.Sprintf("0 %v", d.Units)
	}

	return fmt.Sprintf("%v %v", *d.Value, d.Units)
}

// StringData is a response type in which a string value is stored
type StringData struct {
	Value *string `json:"value"`
}

// String represent the string value of StringData
func (d *StringData) String() string {
	if d.Value == nil {
		return ""
	}

	return *d.Value
}

// StringData is a response type in which a time.Time value is stored
// it's used for marshalling and unmarshalling the json datetime responses
type TimeData struct {
	Value time.Time `json:"value"`
}

// UnmarshalJSON unmarshalls the provided byte slice to a TimeData object
func (t *TimeData) UnmarshalJSON(b []byte) error {
	var tempStruct struct {
		Value string `json:"value"`
	}

	err := json.Unmarshal(b, &tempStruct)

	if err != nil {
		return err
	}

	pt, err := time.Parse(time.RFC3339, tempStruct.Value)

	if err != nil {
		return err
	}

	*t = TimeData{Value: pt}
	return nil
}

// MarshalJson marshals the TimeData object to a byte slice
func (t *TimeData) MarshalJSON() ([]byte, error) {
	return []byte(t.Value.String()), nil
}
