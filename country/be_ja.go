//go:build (lang_ja || lang_all) && (country_all || country_be || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Japanese, "ベルギー")
	dataBelgium.RegisterOfficialName(xlanguage.Japanese, "ベルギー王国")
	dataBelgium.RegisterCapital(xlanguage.Japanese, "ブリュッセル")
}
