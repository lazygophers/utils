//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bz || country_central_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Japanese, "ベリーズ")
	dataBelize.RegisterOfficialName(xlanguage.Japanese, "ベリーズ")
	dataBelize.RegisterCapital(xlanguage.Japanese, "ベルモパン")
}
