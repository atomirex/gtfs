package gtfs

import (
	"encoding/csv"
	"errors"
	"io"
	"reflect"
)

// From agency.txt
type Agency struct {
	Id       string `gtfs_name:"agency_id" gtfs_required:"false"`
	Name     string `gtfs_name:"agency_name" gtfs_required:"true"`
	Url      string `gtfs_name:"agency_url" gtfs_required:"true"`
	Timezone string `gtfs_name:"agency_timezone" gtfs_required:"true"`
	Language string `gtfs_name:"agency_lang" gtfs_required:"false"`
	Phone    string `gtfs_name:"agency_phone" gtfs_required:"false"`
	FareUrl  string `gtfs_name:"agency_fare_url" gtfs_required:"false"`
	Email    string `gtfs_name:"agency_email" gtfs_required:"false"`
}

func (a *Agency) String() string {
	return a.Name + " " + a.Url
}

// From stops.txt
type Stop struct {
	Id                 string `gtfs_name:"stop_id" gtfs_required:"true"`
	Code               string `gtfs_name:"stop_code" gtfs_required:"false"`
	Name               string `gtfs_name:"stop_name" gtfs_required:"true"`
	Description        string `gtfs_name:"stop_desc" gtfs_required:"false"`
	Latitude           string `gtfs_name:"stop_lat" gtfs_required:"true"`
	Longitude          string `gtfs_name:"stop_lon" gtfs_required:"true"`
	ZoneId             string `gtfs_name:"zone_id" gtfs_required:"false"`
	Url                string `gtfs_name:"stop_url" gtfs_required:"false"`
	LocationType       string `gtfs_name:"location_type" gtfs_required:"false"`
	ParentStation      string `gtfs_name:"parent_station" gtfs_required:"false"`
	Timezone           string `gtfs_name:"stop_timezone" gtfs_required:"false"`
	WheelchairBoarding string `gtfs_name:"wheelchair_boarding" gtfs_required:"false"`
}

func (s *Stop) String() string {
	return s.Id + " " + s.Name
}

// From routes.txt
type Route struct {
	Id          string `gtfs_name:"route_id" gtfs_required:"true"`
	AgencyId    string `gtfs_name:"agency_id" gtfs_required:"false"`
	ShortName   string `gtfs_name:"route_short_name" gtfs_required:"true"`
	LongName    string `gtfs_name:"route_long_name" gtfs_required:"true"`
	Description string `gtfs_name:"route_desc" gtfs_required:"false"`
	Type        string `gtfs_name:"route_type" gtfs_required:"true"`
	Url         string `gtfs_name:"route_url" gtfs_required:"false"`
	Color       string `gtfs_name:"route_color" gtfs_required:"false"`
	TextColor   string `gtfs_name:"route_text_color" gtfs_required:"false"`
	SortOrder   string `gtfs_name:"route_sort_order" gtfs_required:"false"`
}

func (r *Route) String() string {
	return r.Id + " " + r.ShortName + " " + r.LongName
}

// From trips.txt
type Trip struct {
	RouteId              string `gtfs_name:"route_id" gtfs_required:"true"`
	ServiceId            string `gtfs_name:"service_id" gtfs_required:"true"`
	Id                   string `gtfs_name:"trip_id" gtfs_required:"true"`
	HeadSign             string `gtfs_name:"trip_headsign" gtfs_required:"false"`
	ShortName            string `gtfs_name:"trip_short_name" gtfs_required:"false"`
	DirectionId          string `gtfs_name:"direction_id" gtfs_required:"false"`
	BlockId              string `gtfs_name:"block_id" gtfs_required:"false"`
	ShapeId              string `gtfs_name:"shape_id" gtfs_required:"false"`
	WheelchairAccessible string `gtfs_name:"wheelchair_accessible" gtfs_required:"false"`
	BikesAllowed         string `gtfs_name:"bikes_allowed" gtfs_required:"false"`
}

func (t *Trip) String() string {
	return t.Id + " " + t.RouteId + " " + t.ServiceId
}

// From stop_times.txt
type StopTime struct {
	TripId            string `gtfs_name:"trip_id" gtfs_required:"true"`
	ArrivalTime       string `gtfs_name:"arrival_time" gtfs_required:"true"`
	DepartureTime     string `gtfs_name:"departure_time" gtfs_required:"true"`
	StopId            string `gtfs_name:"stop_id" gtfs_required:"true"`
	StopSequence      string `gtfs_name:"stop_sequence" gtfs_required:"true"`
	StopHeadSign      string `gtfs_name:"stop_headsign" gtfs_required:"false"`
	PickupType        string `gtfs_name:"pickup_type" gtfs_required:"false"`
	DropOffType       string `gtfs_name:"drop_off_type" gtfs_required:"false"`
	ShapeDistTraveled string `gtfs_name:"shape_dist_traveled" gtfs_required:"false"`
	TimePoint         string `gtfs_name:"timepoint" gtfs_required:"false"`
}

func (st *StopTime) String() string {
	return st.TripId + " " + st.StopId
}

// From calendar.txt
type Service struct {
	ServiceId string `gtfs_name:"service_id" gtfs_required:"true"`
	Monday    string `gtfs_name:"monday" gtfs_required:"true"`
	Tuesday   string `gtfs_name:"tuesday" gtfs_required:"true"`
	Wednesday string `gtfs_name:"wednesday" gtfs_required:"true"`
	Thursday  string `gtfs_name:"thursday" gtfs_required:"true"`
	Friday    string `gtfs_name:"friday" gtfs_required:"true"`
	Saturday  string `gtfs_name:"saturday" gtfs_required:"true"`
	Sunday    string `gtfs_name:"sunday" gtfs_required:"true"`
	StartDate string `gtfs_name:"start_date" gtfs_required:"true"`
	EndDate   string `gtfs_name:"end_date" gtfs_required:"true"`
}

func (s *Service) String() string {
	return s.ServiceId + " from " + s.StartDate + " to " + s.EndDate
}

// From calendar_dates.txt
type ServiceException struct {
	ServiceId     string `gtfs_name:"service_id" gtfs_required:"true"`
	Date          string `gtfs_name:"date" gtfs_required:"true"`
	ExceptionType string `gtfs_name:"exception_type" gtfs_required:"true"`
}

func (s *ServiceException) String() string {
	return s.ServiceId + " on " + s.Date + " exception " + s.ExceptionType
}

// From fare_attributes.txt
type Fare struct {
	FareId           string `gtfs_name:"fare_id" gtfs_required:"true"`
	Price            string `gtfs_name:"price" gtfs_required:"true"`
	CurrencyType     string `gtfs_name:"currency_type" gtfs_required:"true"`
	PaymentMethod    string `gtfs_name:"payment_method" gtfs_required:"true"`
	Transfers        string `gtfs_name:"transfers" gtfs_required:"true"`
	AgencyId         string `gtfs_name:"agency_id" gtfs_required:"false"`
	TransferDuration string `gtfs_name:"transfer_duration" gtfs_required:"false"`
}

func (f *Fare) String() string {
	return f.FareId + " " + f.Price + f.CurrencyType
}

// From fare_rules.txt
type FareRule struct {
	FareId        string `gtfs_name:"fare_id" gtfs_required:"true"`
	RouteId       string `gtfs_name:"route_id" gtfs_required:"false"`
	OriginId      string `gtfs_name:"origin_id" gtfs_required:"false"`
	DestinationId string `gtfs_name:"destination_id" gtfs_required:"false"`
	ContainsId    string `gtfs_name:"contains_id" gtfs_required:"false"`
}

func (f *FareRule) String() string {
	return f.FareId + " " + f.RouteId + " " + f.OriginId + " " + f.DestinationId
}

// From shapes.txt
type ShapePoint struct {
	Id           string `gtfs_name:"shape_id" gtfs_required:"true"`
	PtLatitude   string `gtfs_name:"shape_pt_lat" gtfs_required:"true"`
	PtLongitude  string `gtfs_name:"shape_pt_lon" gtfs_required:"true"`
	PtSequence   string `gtfs_name:"shape_pt_sequence" gtfs_required:"true"`
	DistTraveled string `gtfs_name:"shape_dist_traveled" gtfs_required:"false"`
}

func (s *ShapePoint) String() string {
	return s.Id + " " + s.PtSequence + " " + s.PtLatitude + " " + s.PtLongitude
}

// From frequencies.txt
type Frequency struct {
	TripId      string `gtfs_name:"trip_id" gtfs_required:"true"`
	StartTime   string `gtfs_name:"start_time" gtfs_required:"true"`
	EndTime     string `gtfs_name:"end_time" gtfs_required:"true"`
	HeadwaySecs string `gtfs_name:"headway_secs" gtfs_required:"true"`
	ExactTimes  string `gtfs_name:"exact_times" gtfs_required:"false"`
}

func (f *Frequency) String() string {
	return f.TripId + " " + f.StartTime + " " + f.EndTime
}

// From transfers.txt
type Transfer struct {
	FromStopId          string `gtfs_name:"from_stop_id" gtfs_required:"true"`
	ToStopId            string `gtfs_name:"to_stop_id" gtfs_required:"true"`
	TransferType        string `gtfs_name:"transfer_type" gtfs_required:"true"`
	MinimumTransferTime string `gtfs_name:"min_transfer_time" gtfs_required:"false"`
}

func (t *Transfer) String() string {
	return t.FromStopId + " to " + t.ToStopId + " " + t.TransferType
}

func getFieldIndexForStruct(t reflect.Type, name string) (int, error) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("gtfs_name") == name {
			return i, nil
		}
	}

	return -1, errors.New("Field not found " + name)
}

func Decode(r io.Reader, rowtype interface{}) ([]interface{}, error) {
	c := csv.NewReader(r)
	row, err := c.Read()
	if err != nil {
		return nil, err
	}

	if len(row) == 0 {
		return nil, errors.New("No fields")
	}

	fields := make([]int, len(row))
	t := reflect.TypeOf(rowtype).Elem()

	for i := 0; i < len(row); i++ {
		fields[i], err = getFieldIndexForStruct(t, row[i])
		if err != nil {
			return nil, err
		}
	}

	output := make([]interface{}, 0)

	for {
		row, err := c.Read()
		if err == nil {
			o := reflect.New(t)
			for i := 0; i < len(row); i++ {
				o.Elem().Field(fields[i]).SetString(row[i])
			}
			output = append(output, o.Interface())
		} else if err == io.EOF {
			return output, nil
		} else {
			return nil, err
		}
	}

	// TODO validate required fields are loaded

	return nil, errors.New("Not implemented")
}
