//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Arabic, "أراضي فرنسا الجنوبية والقطبية")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Arabic, "أراضي فرنسا الجنوبية والقطبية")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Arabic, "سان بيير")
}
