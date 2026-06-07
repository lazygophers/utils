//go:build (lang_ja || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Japanese, "ラトビア")
	dataLatvia.RegisterOfficialName(xlanguage.Japanese, "ラトビア共和国")
	dataLatvia.RegisterCapital(xlanguage.Japanese, "リガ")
}
