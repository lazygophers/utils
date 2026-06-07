//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Arabic, "جبل طارق")
	dataGibraltar.RegisterOfficialName(xlanguage.Arabic, "إقليم جبل طارق البريطاني فيما وراء البحار")
	dataGibraltar.RegisterCapital(xlanguage.Arabic, "جبل طارق")
}
