package main

import (
    "fmt"
)

const (
    MEDIATYPE_BOOK = iota
    MEDIATYPE_FILM
)

type MediaType int

type Entry struct {
    Name string
    Date string
    Type MediaType
}

func main() {
    entries := []Entry{
        {Name: "The Wrong Man", Date: "05-01-2026", Type: MEDIATYPE_FILM},
        {Name: "Moby Dick", Date: "08-01-2026", Type: MEDIATYPE_BOOK},
    }

    for _, entry := range entries {
        fmt.Println(entry)
    }
}

func (e Entry) String() string {
    typ := ""
    if e.Type == MEDIATYPE_BOOK {
       typ = "Book"
    } else {
        typ = "Film"
    }

    return typ + " | " + e.Name + ": " + e.Date
}
