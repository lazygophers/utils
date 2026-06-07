//go:build country_africa || country_all || country_eh || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Arabic, "الصحراء الغربية")
	dataWesternSahara.RegisterOfficialName(xlanguage.Arabic, "الجمهورية العربية الصحراوية الديمقراطية")
	dataWesternSahara.RegisterCapital(xlanguage.Arabic, "العيون")
}
