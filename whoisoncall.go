package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	ics "github.com/arran4/golang-ical"
)

func main() {
	var at string

	flag.StringVar(&at, "at", "", "Override when we're querying, currently accepts an HH:MM time")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fatalf("expected exactly 1 pagerduty ical URL argument")
	}

	r, err := http.Get(flag.Arg(0))
	if err != nil {
		fatalf("failed to get the ical file: %v", err)
	}

	if r.StatusCode != 200 {
		fatalf("unexpected status code", r.StatusCode)
	}

	c, err := ics.ParseCalendar(r.Body)
	if err != nil {
		fatalf("could not parse the calendar: %v", err)
	}

	now := time.Now()

	now = possiblyMonkeyWithTime(now, at)

	var current *ics.VEvent

	// this bit is not optimal, but hey:
	for _, ev := range c.Events() {
		start, err := ev.GetStartAt()
		if err != nil {
			fatalf("problem with start time: %v", err)
		}
		end, err := ev.GetEndAt()
		if err != nil {
			fatalf("problem with end time: %v", err)
		}

		if start.Before(now) && now.Before(end) {
			if current != nil {
				fatalf("somehow found > 1 event, not expected")
			}
			current = ev
		}
	}

	if current == nil {
		fatalf("I didn't find a current event, so I guess nobody?")
	}

	p := current.GetProperty(ics.ComponentPropertySummary)
	if p == nil {
		fatalf("current event had no description property")
	}
	fmt.Println(p.Value)
}

func fatalf(s string, args ...interface{}) {
	complainf(s, args...)
	os.Exit(1)
}

func complainf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func possiblyMonkeyWithTime(t time.Time, at string) time.Time {
	if at == "" {
		return t
	}

	timeOfDay, err := time.Parse("15:04", at)
	if err != nil {
		fatalf("could not parse at %s: %v", at, err)
	}

	return time.Date(t.Year(), t.Month(), t.Day(),
		timeOfDay.Hour(), timeOfDay.Minute(),
		0, 0, t.Location())
}
