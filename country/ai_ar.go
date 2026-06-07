//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Arabic, "أنغويلا")
	dataAnguilla.RegisterOfficialName(xlanguage.Arabic, "إقليم أنغويلا البريطاني فيما وراء البحار")
	dataAnguilla.RegisterCapital(xlanguage.Arabic, "ذا فالي")
}
