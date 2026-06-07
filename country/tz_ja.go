//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Japanese, "タンザニア")
	dataTanzania.RegisterOfficialName(xlanguage.Japanese, "タンザニア連合共和国")
	dataTanzania.RegisterCapital(xlanguage.Japanese, "ドドマ")
}
