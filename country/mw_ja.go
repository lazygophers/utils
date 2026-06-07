//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Japanese, "マラウイ")
	dataMalawi.RegisterOfficialName(xlanguage.Japanese, "マラウイ共和国")
	dataMalawi.RegisterCapital(xlanguage.Japanese, "リロングウェ")
}
