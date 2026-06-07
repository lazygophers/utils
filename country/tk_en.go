//go:build country_all || country_oceania || country_polynesia || country_tk

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.English, "Tokelau")
	dataTokelau.RegisterOfficialName(xlanguage.English, "Tokelau")
}
