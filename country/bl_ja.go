//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bl || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Japanese, "サン・バルテルミー島")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Japanese, "サン・バルテルミー島")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Japanese, "グスタヴィア")
}
