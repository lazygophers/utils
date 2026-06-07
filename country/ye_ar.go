package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Arabic, "اليمن")
	dataYemen.RegisterOfficialName(xlanguage.Arabic, "الجمهورية اليمنية")
	dataYemen.RegisterCapital(xlanguage.Arabic, "صنعاء")
}
