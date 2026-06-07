//go:build (lang_ja || lang_all) && (country_africa || country_all || country_sl || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Japanese, "シエラレオネ")
	dataSierraLeone.RegisterOfficialName(xlanguage.Japanese, "シエラレオネ共和国")
	dataSierraLeone.RegisterCapital(xlanguage.Japanese, "フリータウン")
}
