//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Japanese, "ハンガリー")
	dataHungary.RegisterOfficialName(xlanguage.Japanese, "ハンガリー")
	dataHungary.RegisterCapital(xlanguage.Japanese, "ブダペスト")
}
