//go:build (lang_ar || lang_all) && (country_all || country_be || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Arabic, "بلجيكا")
	dataBelgium.RegisterOfficialName(xlanguage.Arabic, "مملكة بلجيكا")
	dataBelgium.RegisterCapital(xlanguage.Arabic, "بروكسل")
}
