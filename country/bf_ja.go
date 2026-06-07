//go:build (lang_ja || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Japanese, "ブルキナファソ")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Japanese, "ブルキナファソ")
	dataBurkinaFaso.RegisterCapital(xlanguage.Japanese, "ワガドゥグー")
}
