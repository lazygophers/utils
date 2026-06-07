//go:build (lang_ja || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Japanese, "チェコ")
	dataCzechia.RegisterOfficialName(xlanguage.Japanese, "チェコ共和国")
	dataCzechia.RegisterCapital(xlanguage.Japanese, "プラハ")
}
