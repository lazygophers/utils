//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Arabic, "الفلبين")
	dataPhilippines.RegisterOfficialName(xlanguage.Arabic, "جمهورية الفلبين")
	dataPhilippines.RegisterCapital(xlanguage.Arabic, "مانيلا")
}
