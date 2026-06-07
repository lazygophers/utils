//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Arabic, "سانت هيلينا وأسينشين وتريستان دا كونا")
	dataSaintHelena.RegisterOfficialName(xlanguage.Arabic, "إقليم سانت هيلينا وأسينشين وتريستان دا كونا البريطاني فيما وراء البحار")
	dataSaintHelena.RegisterCapital(xlanguage.Arabic, "جيمستاون")
}
