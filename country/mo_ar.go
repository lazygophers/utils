//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Arabic, "ماكاو")
	dataMacao.RegisterOfficialName(xlanguage.Arabic, "منطقة ماكاو الإدارية الخاصة بجمهورية الصين الشعبية")
	dataMacao.RegisterCapital(xlanguage.Arabic, "ماكاو")
}
