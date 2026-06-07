//go:build country_all || country_micronesia || country_oceania || country_um

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.English, "United States Minor Outlying Islands")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.English, "United States Minor Outlying Islands")
}
