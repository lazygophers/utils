//go:build (lang_ar || lang_all) && (country_all || country_es || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Arabic, "إسبانيا")
	dataSpain.RegisterOfficialName(xlanguage.Arabic, "مملكة إسبانيا")
	dataSpain.RegisterCapital(xlanguage.Arabic, "مدريد")
}
