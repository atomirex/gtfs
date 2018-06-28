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
func TestStopTimes(t *testing.T) {
	s := `trip_id,arrival_time,departure_time,stop_id,stop_sequence,pickup_type,drop_off_type
AWE1,0:06:10,0:06:10,S1,1,0,0
AWE1,,,S2,2,1,3
AWE1,0:06:20,0:06:30,S3,3,0,0
AWE1,,,S5,4,0,0
AWE1,0:06:45,0:06:45,S6,5,0,0
AWD1,0:06:10,0:06:10,S1,1,0,0
AWD1,,,S2,2,0,0
AWD1,0:06:20,0:06:20,S3,3,0,0
AWD1,,,S4,4,0,0
AWD1,,,S5,5,0,0
AWD1,0:06:45,0:06:45,S6,6,0,0`

	out, err := Decode(strings.NewReader(s), &StopTime{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*StopTime).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestServices(t *testing.T) {
	s := `service_id,monday,tuesday,wednesday,thursday,friday,saturday,sunday,start_date,end_date
WE,0,0,0,0,0,1,1,20060701,20060731
WD,1,1,1,1,1,0,0,20060701,20060731`

	out, err := Decode(strings.NewReader(s), &Service{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Service).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}
func TestServiceExceptions(t *testing.T) {
	s := `service_id,date,exception_type
WD,20060703,2
WE,20060703,1
WD,20060704,2
WE,20060704,1`

	out, err := Decode(strings.NewReader(s), &ServiceException{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*ServiceException).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}
func TestFares(t *testing.T) {
	s := `fare_id,price,currency_type,payment_method,transfers,transfer_duration
1,0.00,USD,0,0,0
2,0.50,USD,0,0,0
3,1.50,USD,0,0,0
4,2.00,USD,0,0,0
5,2.50,USD,0,0,0`

	out, err := Decode(strings.NewReader(s), &Fare{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Fare).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestFareRules(t *testing.T) {
	s := `fare_id,route_id,origin_id,destination_id,contains_id
a,TSW,1,1,
a,TSE,1,1,
a,GRT,1,1,
a,GRJ,1,1,
a,SVJ,1,1,
a,JSV,1,1,
a,GRT,2,4,
a,GRJ,4,2,
b,GRT,3,3,
c,GRT,,,6`

	out, err := Decode(strings.NewReader(s), &FareRule{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*FareRule).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestShapePoints(t *testing.T) {
	s := `shape_id,shape_pt_lat,shape_pt_lon,shape_pt_sequence,shape_dist_traveled
A_shp,37.61956,-122.48161,1,0
A_shp,37.64430,-122.41070,2,6.8310
A_shp,37.65863,-122.30839,3,15.8765`

	out, err := Decode(strings.NewReader(s), &ShapePoint{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*ShapePoint).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestFrequencies(t *testing.T) {
	s := `trip_id,start_time,end_time,headway_secs
AWE1,05:30:00,06:30:00,300
AWE1,06:30:00,20:30:00,180
AWE1,20:30:00,28:00:00,420`

	out, err := Decode(strings.NewReader(s), &Frequency{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Frequency).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}

func TestTransfers(t *testing.T) {
	s := `from_stop_id,to_stop_id,transfer_type,min_transfer_time
S6,S7,2,300
S7,S6,3,
S23,S7,1,`

	out, err := Decode(strings.NewReader(s), &Transfer{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) == 0 {
			t.Errorf("No output")
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Transfer).String())
			}
			t.Errorf("Need to validate the struct")
		}
	}
}
