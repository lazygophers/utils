//go:build (lang_ar || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Arabic, "أروبا")
	dataAruba.RegisterOfficialName(xlanguage.Arabic, "أروبا")
	dataAruba.RegisterCapital(xlanguage.Arabic, "أورانيستاد")
}
