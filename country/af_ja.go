//go:build (lang_ja || lang_all) && (country_af || country_all || country_asia || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Japanese, "アフガニスタン")
	dataAfghanistan.RegisterOfficialName(xlanguage.Japanese, "アフガニスタン・イスラム首長国")
	dataAfghanistan.RegisterCapital(xlanguage.Japanese, "カブール")
}
