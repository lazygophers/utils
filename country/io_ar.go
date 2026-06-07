//go:build (lang_ar || lang_all) && (country_all || country_asia || country_eastern_africa || country_io)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Arabic, "إقليم المحيط الهندي البريطاني")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Arabic, "إقليم المحيط الهندي البريطاني")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Arabic, "دييغو غارسيا")
}
