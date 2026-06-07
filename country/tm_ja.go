//go:build (lang_ja || lang_all) && (country_all || country_asia || country_central_asia || country_tm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Japanese, "トルクメニスタン")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Japanese, "トルクメニスタン")
	dataTurkmenistan.RegisterCapital(xlanguage.Japanese, "アシガバート")
}
