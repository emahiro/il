package hatena

import (
	"bytes"
	"encoding/gob"
)

// Feed represents hatena blog feed in Atom format.
type Feed struct {
	Title   string  `xml:"title"`
	Link    string  `xml:"link"`
	Author  string  `xml:"author"`
	Entries []Entry `xml:"entry"`
}

// Entry represents hatena blog entry.
type Entry struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	ID        string `xml:"id"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
	Author    string `xml:"author"`
	Category  string `xml:"category"`
}

// ToBytes encodes Feed to byte array.
func (f *Feed) ToBytes() ([]byte, error) {
	b := bytes.Buffer{}
	if err := gob.NewEncoder(&b).Encode(f); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// RSS represents hatena blog feed in RSS format.
type RSS struct {
	Title         string `xml:"chennel>title"`
	Link          string `xml:"chennel>link"`
	Description   string `xml:"chennel>description"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	Docs          string `xml:"channel>docs"`
	Items         []Item `xml:"item"` // RSS だと ここがうまく Decode されない、ということがわかった
}

// Item represents hatena blog rss item.
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Category    string `xml:"category"`
}

// ToBytes encodes RSS to byte array.
func (r *RSS) ToBytes() ([]byte, error) {
	b := bytes.Buffer{}
	if err := gob.NewEncoder(&b).Encode(r); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
