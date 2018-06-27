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
