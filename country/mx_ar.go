//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Arabic, "المكسيك")
	dataMexico.RegisterOfficialName(xlanguage.Arabic, "الولايات المتحدة المكسيكية")
	dataMexico.RegisterCapital(xlanguage.Arabic, "مدينة مكسيكو")
}
