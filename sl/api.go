package sl

import (
	"fmt"

	"github.com/frozzare/go-util/httputil"
	"github.com/frozzare/slnotify/config"
)

const ENDPOINT = "http://api.sl.se/api2/"

type Response struct {
	ResponseData ResponseData
}

type ResponseData struct {
	StopPointDeviations []StopPointDeviations
}

type StopPointDeviations struct {
	Deviation Deviation
	StopInfo  StopInfo
}

type Deviation struct {
	Text            string
	ImportanceLevel int
	Consequence     string
}

type StopInfo struct {
	StopAreaName  string
	TransportMode string
}

func getEndpoint(path string, siteID int) string {
	c := config.Get()

	return fmt.Sprintf("%s%s.json?siteid=%d&timewindow=%d&key=%s", ENDPOINT, path, siteID, c.SL.TimeWindow, c.SL.APIKey)
}

// GetDeviations returns a list of deviations.
func GetDeviations(siteID int) ([]StopPointDeviations, error) {
	res := new(Response)

	if err := httputil.GetJSON(getEndpoint("realtimedepartures", siteID), &res); err != nil {
		return []StopPointDeviations{}, err
	}

	return res.ResponseData.StopPointDeviations, nil
}
