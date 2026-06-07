//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ml || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Japanese, "マリ共和国")
	dataMali.RegisterOfficialName(xlanguage.Japanese, "マリ共和国")
	dataMali.RegisterCapital(xlanguage.Japanese, "バマコ")
}
