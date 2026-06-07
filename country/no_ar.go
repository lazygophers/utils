//go:build (lang_ar || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Arabic, "النرويج")
	dataNorway.RegisterOfficialName(xlanguage.Arabic, "مملكة النرويج")
	dataNorway.RegisterCapital(xlanguage.Arabic, "أوسلو")
}
