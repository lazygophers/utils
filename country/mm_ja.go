//go:build (lang_ja || lang_all) && (country_all || country_asia || country_mm || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Japanese, "ミャンマー")
	dataMyanmar.RegisterOfficialName(xlanguage.Japanese, "ミャンマー連邦共和国")
	dataMyanmar.RegisterCapital(xlanguage.Japanese, "ネピドー")
}
