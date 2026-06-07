//go:build (lang_ja || lang_all) && (country_all || country_europe || country_pt || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Japanese, "ポルトガル")
	dataPortugal.RegisterOfficialName(xlanguage.Japanese, "ポルトガル共和国")
	dataPortugal.RegisterCapital(xlanguage.Japanese, "リスボン")
}
