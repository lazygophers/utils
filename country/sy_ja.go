//go:build (lang_ja || lang_all) && (country_all || country_asia || country_sy || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Japanese, "シリア")
	dataSyria.RegisterOfficialName(xlanguage.Japanese, "シリア・アラブ共和国")
	dataSyria.RegisterCapital(xlanguage.Japanese, "ダマスカス")
}
