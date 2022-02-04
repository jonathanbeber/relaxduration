package relaxduration

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type RelaxedDuration struct {
	d *time.Duration
}

func (r RelaxedDuration) String() string {
	if r.d != nil {
		return r.d.String()
	}
	return ""
}

func (r RelaxedDuration) Set(s string) error {
	if d, err := Parse(s); err != nil {
		return err
	} else {
		*r.d = d
	}

	return nil
}

func Parse(d string) (time.Duration, error) {
	d = strings.ReplaceAll(d, " ", "")
	d = strings.ReplaceAll(d, "and", "")
	d = strings.ReplaceAll(d, "with", "")

	if d == "" {
		return time.Duration(0), nil
	}

	expiration, err := time.ParseDuration(d)
	if err == nil {
		return expiration, nil
	}

	for i, v := range d {
		if unicode.IsDigit(v) {
			continue
		}
		magnitude, err := strconv.Atoi(d[0:i])

		if err != nil {
			return expiration, errors.New("invalid format")
		}

		unit, remaining := getUnitAndRemaining(d, i)
		var duration time.Duration
		switch unit {
		case "s", "sec", "secs", "second", "seconds":
			duration = time.Second * time.Duration(magnitude)
		case "m", "min", "mins", "minute", "minutes":
			duration = time.Minute * time.Duration(magnitude)
		case "h", "hour", "hours":
			duration = time.Hour * time.Duration(magnitude)
		case "d", "day", "days":
			duration = time.Hour * 24 * time.Duration(magnitude)
		case "w", "week", "weeks":
			duration = time.Hour * 24 * 7 * time.Duration(magnitude)
		default:
			return expiration, errors.New("invalid format")
		}

		d, err := Parse(remaining)
		if err != nil {
			return expiration, err
		}
		return duration + d, nil
	}
	return expiration, errors.New("invalid format")
}

func getUnitAndRemaining(d string, i int) (string, string) {
	var j int
	for _, l := range d[i+1:] {
		if unicode.IsDigit(l) {
			break
		}
		j++
	}
	return d[i : i+j+1], d[i+j+1:]
}
