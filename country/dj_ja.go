//go:build (lang_ja || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Japanese, "ジブチ")
	dataDjibouti.RegisterOfficialName(xlanguage.Japanese, "ジブチ共和国")
	dataDjibouti.RegisterCapital(xlanguage.Japanese, "ジブチ市")
}
