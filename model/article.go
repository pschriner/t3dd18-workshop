package model

import (
	"fmt"
	"time"
)

// Article is a CMS article
type Article struct {
	UID       int64
	Title     string
	Text      string
	Image     Image
	Tstamp    *time.Time
	Startdate *time.Time
	Enddate   *time.Time
}

// IsVisible checks whether this article should be visible
func (a Article) IsVisible() bool {
	return (a.Startdate == nil || a.Startdate.Before(time.Now())) && (a.Enddate == nil || a.Enddate.After(time.Now()))
}

// Tostring returns a string representation of the article
func (a Article) Tostring() string {
	return fmt.Sprintf("Title: %s\nText: %s", a.Title, a.Text)
}
