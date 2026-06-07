//go:build (lang_ja || lang_all) && (country_africa || country_all || country_sn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Japanese, "セネガル")
	dataSenegal.RegisterOfficialName(xlanguage.Japanese, "セネガル共和国")
	dataSenegal.RegisterCapital(xlanguage.Japanese, "ダカール")
}
