//go:build (lang_ar || lang_all) && (country_all || country_asia || country_eastern_asia || country_mn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Arabic, "منغوليا")
	dataMongolia.RegisterOfficialName(xlanguage.Arabic, "منغوليا")
	dataMongolia.RegisterCapital(xlanguage.Arabic, "أولان باتور")
}
