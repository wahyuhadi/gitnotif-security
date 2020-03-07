package github

import "time"

// For parsing time => get 1 month since
func TimeParse() time.Time {
	now := time.Now()
	then := now.AddDate(0, 0, -1).Format("2006-01-02T15:04:05:0700")
	since, _ := time.Parse("2006-01-02", then)
	return since
}
