//go:build country_all || country_bg || country_eastern_europe || country_europe

package fake

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

func init() {
	c := country.Get("BG")
	register(&Locale{
		Country:        c,
		OfficialLangs:  c.SpokenLanguages(),
		PhonePrefixes:  nil,
		LandlinePrefix: nil,
		ZipFormat:      "",
		IdCardGen:      nil,
		Streets:        map[xlanguage.Tag][]string{},
		Cities:         map[xlanguage.Tag][]CityEntry{},
		FirstNames:     map[xlanguage.Tag]map[Gender][]string{},
		LastNames:      map[xlanguage.Tag][]string{},
		Domain:         "bg",
	})
}
