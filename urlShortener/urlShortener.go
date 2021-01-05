package urlShortener

import (
	"fmt"
	"time"
)

type Shortener struct {
	urls   map[string]string
	visits map[string]uint64
}

type ShortenerError struct {
	When time.Time
	What string
}

var instance Shortener

////////////////////////////////////////////////////////////////////////
// Exposed Methods
////////////////////////////////////////////////////////////////////////

/**
Method to initialize the Shortener
*/
func InitShortener() Shortener {
	instance = Shortener{urls: make(map[string]string), visits: make(map[string]uint64)}
	return instance
}

/**
Adds a url to the urls map using an alias
*/
func (s *Shortener) AddUrl(alias, url string) error {
	// Check if the alias already exists
	if _, ok := s.urls[alias]; ok {
		return &ShortenerError{
			time.Now(),
			"Alias already exists",
		}
	}

	// Add to dictionary and to visits map
	s.urls[alias] = url
	s.visits[alias] = 0

	return nil
}

/**
Removes a given alias from the map
*/
func (s *Shortener) RemoveUrl(alias string) (string, error) {
	if _, ok := s.urls[alias]; !ok {
		return "", &ShortenerError{
			time.Now(),
			"Alias does not exist",
		}
	}

	val := s.urls[alias]

	// Add to dictionary and to visits map
	delete(s.urls, alias)
	delete(s.visits, alias)

	return val, nil
}

func (s *Shortener) GetUrl(alias string) (string, error) {
	if _, ok := s.urls[alias]; !ok {
		return "", &ShortenerError{
			time.Now(),
			"Alias does not exist",
		}
	}

	// Add a visit to the url
	s.visits[alias] = s.visits[alias] + 1

	return s.urls[alias], nil
}

func (s *Shortener) GetUrlVisits(alias string) (uint64, error) {
	if _, ok := s.urls[alias]; !ok {
		return 0, &ShortenerError{
			time.Now(),
			"Alias does not exist",
		}
	}

	return s.visits[alias], nil
}

func (s *Shortener) GetGeneralVisits() map[string]uint64 {
	return s.visits
}

func (s *Shortener) GetGeneralUrls() map[string]string {
	return s.urls
}

func (e *ShortenerError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

////////////////////////////////////////////////////////////////////////
// Internal Methods
////////////////////////////////////////////////////////////////////////
