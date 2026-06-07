//go:build country_all || country_europe || country_je || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.English, "Jersey")
	dataJersey.RegisterOfficialName(xlanguage.English, "Bailiwick of Jersey")
	dataJersey.RegisterCapital(xlanguage.English, "Saint Helier")
}
