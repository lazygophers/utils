//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Arabic, "الولايات المتحدة")
	dataUnitedStates.RegisterOfficialName(xlanguage.Arabic, "الولايات المتحدة الأمريكية")
	dataUnitedStates.RegisterCapital(xlanguage.Arabic, "واشنطن العاصمة")
}
