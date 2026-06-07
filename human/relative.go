package human

import (
	"fmt"
	"time"
)

// RelativeTime renders the gap between t and now as a localized phrase
// such as "3 minutes ago" / "in 2 hours".
func RelativeTime(t time.Time) string { return formatRelativeTime(t) }

func formatRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	locale, _ := GetLocaleConfig(currentTag())

	if diff < 0 {
		diff = -diff
		return formatFutureTime(diff, locale)
	}

	return formatPastTime(diff, locale)
}

func formatPastTime(diff time.Duration, locale *Locale) string {
	if diff < 10*time.Second {
		return locale.RelativeTime.JustNow
	}

	if diff < time.Minute {
		seconds := int(diff.Seconds())
		return fmt.Sprintf(locale.RelativeTime.SecondsAgo, seconds)
	}

	if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf(locale.RelativeTime.MinutesAgo, minutes)
	}

	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf(locale.RelativeTime.HoursAgo, hours)
	}

	if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf(locale.RelativeTime.DaysAgo, days)
	}

	if diff < 30*24*time.Hour {
		weeks := int(diff.Hours() / (7 * 24))
		return fmt.Sprintf(locale.RelativeTime.WeeksAgo, weeks)
	}

	if diff < 365*24*time.Hour {
		months := int(diff.Hours() / (30 * 24))
		return fmt.Sprintf(locale.RelativeTime.MonthsAgo, months)
	}

	years := int(diff.Hours() / (365 * 24))
	return fmt.Sprintf(locale.RelativeTime.YearsAgo, years)
}

func formatFutureTime(diff time.Duration, locale *Locale) string {
	if diff < time.Minute {
		seconds := int(diff.Seconds())
		return fmt.Sprintf(locale.RelativeTime.SecondsLater, seconds)
	}

	if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf(locale.RelativeTime.MinutesLater, minutes)
	}

	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf(locale.RelativeTime.HoursLater, hours)
	}

	if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf(locale.RelativeTime.DaysLater, days)
	}

	if diff < 30*24*time.Hour {
		weeks := int(diff.Hours() / (7 * 24))
		return fmt.Sprintf(locale.RelativeTime.WeeksLater, weeks)
	}

	if diff < 365*24*time.Hour {
		months := int(diff.Hours() / (30 * 24))
		return fmt.Sprintf(locale.RelativeTime.MonthsLater, months)
	}

	years := int(diff.Hours() / (365 * 24))
	return fmt.Sprintf(locale.RelativeTime.YearsLater, years)
}
