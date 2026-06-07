//go:build (lang_ja || lang_all) && (country_africa || country_all || country_bj || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Japanese, "ベナン")
	dataBenin.RegisterOfficialName(xlanguage.Japanese, "ベナン共和国")
	dataBenin.RegisterCapital(xlanguage.Japanese, "ポルトノボ")
}
