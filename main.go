package main

import (
	"flag"
	"fmt"

	"strings"

	"github.com/frozzare/go-util/pathutil"
	"github.com/frozzare/go-util/sliceutil"
	"github.com/frozzare/sl/config"
	"github.com/frozzare/sl/notify"
	"github.com/frozzare/sl/sl"
)

var (
	configFlag = flag.String("config", "config.json", "Path to config file")
	siteIDFlag = flag.Int("site-id", 0, "sl.se site id")
)

func main() {
	flag.Parse()

	// Init config.
	config.Init(pathutil.RealPath(*configFlag))
	c := config.Get()

	if *siteIDFlag == 0 {
		fmt.Println("Site id is empty")
		return
	}

	// Get deviations from sl.se for a site id.
	deviations, err := sl.GetDeviations(*siteIDFlag)

	if len(deviations) == 0 {
		fmt.Println("No result from sl.se")
		return
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return
	}

	text := ""
	name := deviations[0].StopInfo.StopAreaName

	for _, d := range deviations {
		if d.Deviation.Text == "" {
			continue
		}

		if sliceutil.Has(c.SL.TransportMode, strings.ToLower(d.StopInfo.TransportMode)) {
			continue
		}

		text += d.Deviation.Text + "\n\n"
	}

	if text == "" {
		text = "Inga avvikelser"
	} else {
		text = text[0 : len(text)-2]
	}

	err = notify.Push(fmt.Sprintf("%s station", name), text)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		return
	}

	fmt.Println("Sent notification to Pushover!")
}
