//go:build (lang_ar || lang_all) && (country_all || country_by || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Arabic, "بيلاروس")
	dataBelarus.RegisterOfficialName(xlanguage.Arabic, "جمهورية بيلاروس")
	dataBelarus.RegisterCapital(xlanguage.Arabic, "مينسك")
}
