package main

import (
	"fmt"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func ParseStringToTime(dateString, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	diff := now.Sub(pastTime)

	if diff < time.Minute {
		seconds := int(diff.Seconds())
		if seconds == 1 {
			return "1 second ago"
		}
		return fmt.Sprintf("%d seconds ago", seconds)
	}

	if diff < time.Hour {
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	days := int(diff.Hours() / 24)
	if days == 1 {
		return "1 day ago"
	}
	if days < 30 {
		return fmt.Sprintf("%d days ago", days)
	}

	months := int(days / 30)
	if months == 1 {
		return "1 month ago"
	}
	if months < 12 {
		return fmt.Sprintf("%d months ago", months)
	}

	years := int(months / 12)
	if years == 1 {
		return "1 year ago"
	}
	return fmt.Sprintf("%d years ago", years)
}

func NextWorkday(start time.Time) time.Time {
	return start.Round(24 * time.Hour * 7)
}

func main() {
	fmt.Println(NextWorkday(time.Now()))
}
