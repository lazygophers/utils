//go:build country_africa || country_all || country_northern_africa || country_tn

package fake

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

func init() {
	c := country.Get("TN")
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
		Domain:         "tn",
	})
}
