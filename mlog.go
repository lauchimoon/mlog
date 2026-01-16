package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "sort"
    "time"
    "unicode"
)

type Entry struct {
    Name     string
    Date     string
    DateUnix int64
    Type     string
}

const (
    ProgramName = "mlog"
    LogFileName = "log.txt"
)

func main() {
    logFile, err := os.OpenFile(LogFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        fmt.Printf("%s: %v\n", ProgramName, err)
        return
    }
    defer logFile.Close()

    addMedia := flag.Bool("add", false, "Create new entry in log")
    listByType := flag.String("lstype", "", "List entries by type. Empty value is equivalent to listing all entries")
    flag.Parse()

    entries := []Entry{}
    scanner := bufio.NewScanner(logFile)

    for scanner.Scan() {
        entries = append(entries, NewEntryFromString(scanner.Text()))
    }

    sort.Slice(entries, func(i, j int) bool {
        if entries[i].DateUnix != entries[j].DateUnix {
            return entries[i].DateUnix < entries[j].DateUnix
        }

        return entries[i].Name < entries[j].Name
    })

    if *addMedia {
        args := flag.Args()
        ent := Entry{
            Name: args[1],
            Date: args[2],
            Type: args[0],
        }

        _, err = logFile.WriteString(fmt.Sprintf("%s|%s|%s\n", strings.ToLower(ent.Type), ent.Name, ent.Date))
        if err != nil {
            fmt.Printf("%s: %v\n", ProgramName, err)
            return
        }
    }

    if *listByType == "" {
        for _, entry := range entries {
            fmt.Println(entry)
        }
    } else {
        for _, entry := range entries {
            if entry.Type == strings.ToLower(*listByType) {
                fmt.Println(entry)
            }
        }
    }
}

func NewEntryFromString(entryString string) Entry {
    parts := strings.Split(entryString, "|")
    for i := range parts {
        parts[i] = strings.TrimSpace(parts[i])
    }

    ent := Entry{Name: parts[1], Type: parts[0]}
    dateString := parts[2]
    date, err := time.Parse(time.DateTime, dateString)
    if err != nil {
        date, err = time.Parse(time.DateOnly, dateString)
        if err != nil {
            return Entry{}
        }
    }

    ent.Date = dateString
    ent.DateUnix = date.Unix()
    return ent
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
