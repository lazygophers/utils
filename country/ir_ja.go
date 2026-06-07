//go:build (lang_ja || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Japanese, "イラン")
	dataIran.RegisterOfficialName(xlanguage.Japanese, "イラン・イスラム共和国")
	dataIran.RegisterCapital(xlanguage.Japanese, "テヘラン")
}
