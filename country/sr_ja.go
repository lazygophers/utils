//go:build (lang_ja || lang_all) && (country_all || country_americas || country_south_america || country_sr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Japanese, "スリナム")
	dataSuriname.RegisterOfficialName(xlanguage.Japanese, "スリナム共和国")
	dataSuriname.RegisterCapital(xlanguage.Japanese, "パラマリボ")
}
