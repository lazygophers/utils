//go:build (lang_ja || lang_all) && (country_africa || country_all || country_gw || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Japanese, "ギニアビサウ")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Japanese, "ギニアビサウ共和国")
	dataGuineaBissau.RegisterCapital(xlanguage.Japanese, "ビサウ")
}
