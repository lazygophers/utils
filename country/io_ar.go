//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Arabic, "إقليم المحيط الهندي البريطاني")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Arabic, "إقليم المحيط الهندي البريطاني")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Arabic, "دييغو غارسيا")
}
