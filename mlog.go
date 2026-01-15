package main

import (
    "fmt"
    "strings"
    "time"
)

const (
    MediatypeBook = iota
    MediatypeFilm
)

type MediaType int

type Entry struct {
    Name string
    Date string
    Type MediaType
}

func main() {
    entries := []Entry{
        {Name: "The Wrong Man", Date: "2026-01-05", Type: MediatypeFilm},
        {Name: "Moby Dick", Date: "2026-01-08", Type: MediatypeBook},
        {Name: "Chungking Express", Date: "2026-01-12 09:40:00", Type: MediatypeFilm},
    }

    for _, entry := range entries {
        fmt.Println(entry)
    }
}

func (ent Entry) String() string {
    typ := ""
    if ent.Type == MediatypeBook {
       typ = "Book"
    } else {
        typ = "Film"
    }

    dateFields := strings.Fields(ent.Date)
    return typ + " | " + ent.Name + ": " + dateFields[0] + " (" + getTimeSince(ent.Date) + ")"
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
