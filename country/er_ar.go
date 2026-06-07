//go:build country_africa || country_all || country_eastern_africa || country_er

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Arabic, "إريتريا")
	dataEritrea.RegisterOfficialName(xlanguage.Arabic, "دولة إريتريا")
	dataEritrea.RegisterCapital(xlanguage.Arabic, "أسمرة")
}
