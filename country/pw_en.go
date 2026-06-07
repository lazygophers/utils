//go:build country_all || country_micronesia || country_oceania || country_pw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.English, "Palau")
	dataPalau.RegisterOfficialName(xlanguage.English, "Republic of Palau")
	dataPalau.RegisterCapital(xlanguage.English, "Ngerulmud")
}
