//go:build (lang_ja || lang_all) && (country_all || country_asia || country_central_asia || country_uz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Japanese, "ウズベキスタン")
	dataUzbekistan.RegisterOfficialName(xlanguage.Japanese, "ウズベキスタン共和国")
	dataUzbekistan.RegisterCapital(xlanguage.Japanese, "タシュケント")
}
