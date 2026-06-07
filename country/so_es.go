//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_so)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Spanish, "Somalia")
	dataSomalia.RegisterOfficialName(xlanguage.Spanish, "República Federal de Somalia")
	dataSomalia.RegisterCapital(xlanguage.Spanish, "Mogadiscio")
}
