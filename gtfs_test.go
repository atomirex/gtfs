package gtfs

import (
	"strings"
	"testing"
)

func assert(t *testing.T, condition bool, message string) {
	if !condition {
		t.Errorf(message)
	}
}

func TestAgency(t *testing.T) {
	s := `agency_id,agency_name,agency_url,agency_timezone,agency_phone,agency_lang
FunBus,The Fun Bus,http://www.thefunbus.org,America/Los_Angeles,(310) 555-0222,en`

	out, err := Decode(strings.NewReader(s), &Agency{})
	if err != nil {
		t.Error(err)
	} else {
		if len(out) != 1 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			a := out[0].(*Agency)
			t.Log(a.String())

			assert(t, a.Id == "FunBus", "Wrong agency Id")
			assert(t, a.Name == "The Fun Bus", "Wrong agency name")
			assert(t, a.Url == "http://www.thefunbus.org", "Wrong agency url")
			assert(t, a.Timezone == "America/Los_Angeles", "Wrong agency timezone")
			assert(t, a.Phone == "(310) 555-0222", "Wrong agency phone number")
			assert(t, a.Language == "en", "Wrong agency language")
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
		if len(out) != 8 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Stop).String())
			}

			s := out[5].(*Stop)

			assert(t, s.Id == "S6", "Wrong stop ID")
			assert(t, s.Name == "Mission St. & 15th St.", "Wrong stop name")
			assert(t, s.Description == "The stop is located 10 feet north of Mission St.", "Wrong stop description")
			assert(t, s.Latitude == "37.766629", "Wrong stop latitude")
			assert(t, s.Longitude == "-122.419782", "Wrong stop longitude")
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
		if len(out) != 1 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Route).String())
			}

			r := out[0].(*Route)

			assert(t, r.Id == "A", "Wrong route id")
			assert(t, r.ShortName == "17", "Wrong route short name")
			assert(t, r.LongName == "Mission", "Wrong long name for route")
			assert(t, r.Description == `The "A" route travels from lower Mission to Downtown.`, "Wrong route description")
			assert(t, r.Type == "3", "Wrong route type")
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
		if len(out) != 2 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Trip).String())
			}

			r := out[0].(*Trip)

			assert(t, r.RouteId == "A", "Wrong trip route id")
			assert(t, r.ServiceId == "WE", "Wrong trip service id")
			assert(t, r.Id == "AWE1", "Wrong trip id")
			assert(t, r.HeadSign == "Downtown", "Wrong trip headsign")
			assert(t, r.BlockId == "1", "Wrong trip block id")
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
		if len(out) != 11 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*StopTime).String())
			}

			s := out[10].(*StopTime)

			assert(t, s.TripId == "AWD1", "Wrong stop time trip id")
			assert(t, s.ArrivalTime == "0:06:45", "Wrong stop time arrival time")
			assert(t, s.DepartureTime == "0:06:45", "Wrong stop time departure time")
			assert(t, s.StopId == "S6", "Wrong stop time stop id")
			assert(t, s.StopSequence == "6", "Wrong stop time stop sequence")
			assert(t, s.PickupType == "0", "Wrong stop time pickup type")
			assert(t, s.DropOffType == "0", "Wrong stop time drop off type")
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
		if len(out) != 2 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Service).String())
			}

			s := out[0].(*Service)

			assert(t, s.ServiceId == "WE", "Wrong service id")
			assert(t, s.Monday == "0", "Wrong service monday")
			assert(t, s.Tuesday == "0", "Wrong service tuesday")
			assert(t, s.Wednesday == "0", "Wrong service wednesday")
			assert(t, s.Thursday == "0", "Wrong service thursday")
			assert(t, s.Friday == "0", "Wrong service friday")
			assert(t, s.Saturday == "1", "Wrong service saturday")
			assert(t, s.Sunday == "1", "Wrong service sunday")
			assert(t, s.StartDate == "20060701", "Wrong service start date")
			assert(t, s.EndDate == "20060731", "Wrong service end date")
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
		if len(out) != 4 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*ServiceException).String())
			}

			s := out[2].(*ServiceException)

			assert(t, s.ServiceId == "WD", "Wrong service exception service id")
			assert(t, s.Date == "20060704", "Wrong service exception date")
			assert(t, s.ExceptionType == "2", "Wrong service exception type")
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
		if len(out) != 5 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Fare).String())
			}

			f := out[2].(*Fare)

			assert(t, f.FareId == "3", "Wrong fare id")
			assert(t, f.Price == "1.50", "Wrong fare price")
			assert(t, f.CurrencyType == "USD", "Wrong fare currency type")
			assert(t, f.PaymentMethod == "0", "Wrong fare payment method")
			assert(t, f.Transfers == "0", "Wrong fare transfers")
			assert(t, f.TransferDuration == "0", "Wrong fare transfer duration")
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
		if len(out) != 10 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*FareRule).String())
			}

			f := out[2].(*FareRule)

			assert(t, f.FareId == "a", "Wrong fare rule fare id")
			assert(t, f.RouteId == "GRT", "Wrong fare rules route id")
			assert(t, f.OriginId == "1", "Wrong fare rule origin id")
			assert(t, f.DestinationId == "1", "Wrong fare rule destination id")
			assert(t, f.ContainsId == "", "Wrong fare rules contains id")
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
		if len(out) != 3 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*ShapePoint).String())
			}

			sp := out[0].(*ShapePoint)

			assert(t, sp.Id == "A_shp", "Wrong shape point shape id")
			assert(t, sp.PtLatitude == "37.61956", "Wrong shape points latitude")
			assert(t, sp.PtLongitude == "-122.48161", "Wrong shape points longitude")
			assert(t, sp.PtSequence == "1", "Wrong shape points sequence")
			assert(t, sp.DistTraveled == "0", "Wrong shape points distance travelled")
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
		if len(out) != 3 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Frequency).String())
			}

			f := out[1].(*Frequency)

			assert(t, f.TripId == "AWE1", "Wrong frequency trid id")
			assert(t, f.StartTime == "06:30:00", "Wrong frequency start time")
			assert(t, f.EndTime == "20:30:00", "Wrong frequency end time")
			assert(t, f.HeadwaySecs == "180", "Wrong frequency headway")
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
		if len(out) != 3 {
			t.Errorf("Wrong length of output %d", len(out))
		} else {
			for i := 0; i < len(out); i++ {
				t.Log(out[i].(*Transfer).String())
			}

			tr := out[0].(*Transfer)
			assert(t, tr.FromStopId == "S6", "Wrong transfer from stop id")
			assert(t, tr.ToStopId == "S7", "Wrong transfer to stop id")
			assert(t, tr.TransferType == "2", "Wrong transfer type")
			assert(t, tr.MinimumTransferTime == "300", "Wrong transfer min transfer time")
		}
	}
}
