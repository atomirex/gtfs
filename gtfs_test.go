package gtfs

import (
	"strings"
	"testing"
)

func TestAgency(t *testing.T) {
	s := `agency_id,agency_name,agency_url,agency_timezone,agency_phone,agency_lang
FunBus,The Fun Bus,http://www.thefunbus.org,America/Los_Angeles,(310) 555-0222,en`

	out, err := Decode(strings.NewReader(s), &Agency{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			t.Log(out[0].(*Agency).String())
			t.Errorf("Need to validate the struct")
		}
	}
}
func TestStops(t *testing.T) {
	s := `stop_id,stop_name,stop_desc,stop_lat,stop_lon,stop_url,location_type,parent_station
S1,Mission St. & Silver Ave.,The stop is located at the southwest corner of the intersection.,37.728631,-122.431282,,,
S2,Mission St. & Cortland Ave.,The stop is located 20 feet south of Mission St.,37.74103,-122.422482,,,
S3,Mission St. & 24th St.,The stop is located at the southwest corner of the intersection.,37.75223,-122.418581,,,
S4,Mission St. & 21st St.,The stop is located at the northwest corner of the intersection.,37.75713,-122.418982,,,
S5,Mission St. & 18th St.,The stop is located 25 feet west of 18th St.,37.761829,-122.419382,,,
S6,Mission St. & 15th St.,The stop is located 10 feet north of Mission St.,37.766629,-122.419782,,,
S7,24th St. Mission Station,,37.752240,-122.418450,,,S8
S8,24th St. Mission Station,,37.752240,-122.418450,http://www.bart.gov/stations/stationguide/stationoverview_24st.asp,1,`

	out, err := Decode(strings.NewReader(s), &Stop{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Stop).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestRoutes(t *testing.T) {
	s := `route_id,route_short_name,route_long_name,route_desc,route_type
A,17,Mission,"The ""A"" route travels from lower Mission to Downtown.",3`

	out, err := Decode(strings.NewReader(s), &Route{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Route).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}
func TestTrips(t *testing.T) {
	s := `route_id,service_id,trip_id,trip_headsign,block_id
A,WE,AWE1,Downtown,1
A,WE,AWE2,Downtown,2`

	out, err := Decode(strings.NewReader(s), &Trip{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Trip).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}
