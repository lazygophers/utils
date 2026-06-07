//go:build (lang_ja || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Japanese, "パキスタン")
	dataPakistan.RegisterOfficialName(xlanguage.Japanese, "パキスタン・イスラム共和国")
	dataPakistan.RegisterCapital(xlanguage.Japanese, "イスラマバード")
}
