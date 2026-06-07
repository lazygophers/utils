//go:build (lang_ja || lang_all) && (country_all || country_europe || country_it || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Japanese, "イタリア")
	dataItaly.RegisterOfficialName(xlanguage.Japanese, "イタリア共和国")
	dataItaly.RegisterCapital(xlanguage.Japanese, "ローマ")
}
