//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Japanese, "ウガンダ")
	dataUganda.RegisterOfficialName(xlanguage.Japanese, "ウガンダ共和国")
	dataUganda.RegisterCapital(xlanguage.Japanese, "カンパラ")
}
