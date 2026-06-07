//go:build country_all || country_gu || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.English, "Guam")
	dataGuam.RegisterOfficialName(xlanguage.English, "Territory of Guam")
	dataGuam.RegisterCapital(xlanguage.English, "Hagatna")
}
