//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ly || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Japanese, "リビア")
	dataLibya.RegisterOfficialName(xlanguage.Japanese, "リビア国")
	dataLibya.RegisterCapital(xlanguage.Japanese, "トリポリ")
}
