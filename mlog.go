package main

import (
    "fmt"
    "strings"
    "time"
    "unicode"
)

type Entry struct {
    Name string
    Date string
    Type string
}

func main() {
    entries := []Entry{
        NewEntryFromString("film         |    The Wrong Man|2026-01-05  "),
        {Name: "Moby Dick", Date: "2026-01-08", Type: "book"},
        {Name: "Chungking Express", Date: "2026-01-12 09:40:00", Type: "FiLm"},
    }

    for _, entry := range entries {
        fmt.Println(entry)
    }
}

func NewEntryFromString(entryString string) Entry {
    parts := strings.Split(entryString, "|")
    for i := range parts {
        parts[i] = strings.TrimSpace(parts[i])
    }

    return Entry{Name: parts[1], Date: parts[2], Type: parts[0]}
}

func (ent Entry) String() string {
    dateFields := strings.Fields(ent.Date)
    return capitalize(ent.Type) + " | " + ent.Name + ": " + dateFields[0] + " (" + getTimeSince(ent.Date) + ")"
}

func capitalize(s string) string {
    builder := strings.Builder{}
    firstChar := rune(s[0])
    builder.WriteRune(unicode.ToUpper(firstChar))

    for _, c := range s[1:] {
        builder.WriteRune(unicode.ToLower(c))
    }

    return builder.String()
}

func getTimeSince(dateString string) string {
    date, err := time.Parse(time.DateTime, dateString)
    if err != nil {
        date, err = time.Parse(time.DateOnly, dateString)
        if err != nil {
            return ""
        }
    }

    duration := time.Since(date)

    hours := duration.Hours()
    minutes := duration.Minutes()
    seconds := duration.Seconds()

    years := int64(hours/8760.0)
    months := int64(hours * 0.001369)
    days := int64(hours / 24.0)

    agoText := ""
    var metric int64 = 0

    if years > 0 {
        agoText = "years ago"
        metric = years
    } else if months > 0 {
        agoText = "months ago"
        metric = months
    } else if days > 0 {
        agoText = "days ago"
        metric = days
    } else if hours > 0 {
        agoText = "hours ago"
        metric = int64(hours)
    } else if minutes > 0 {
        agoText = "minutes ago"
        metric = int64(minutes)
    } else {
        agoText = "seconds ago"
        metric = int64(seconds)
    }

    return fmt.Sprintf("%v %s", metric, agoText)
}
