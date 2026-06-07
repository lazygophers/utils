//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Japanese, "ウクライナ")
	dataUkraine.RegisterOfficialName(xlanguage.Japanese, "ウクライナ")
	dataUkraine.RegisterCapital(xlanguage.Japanese, "キーウ")
}
