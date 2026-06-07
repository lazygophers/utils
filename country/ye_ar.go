//go:build country_all || country_asia || country_western_asia || country_ye

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Arabic, "اليمن")
	dataYemen.RegisterOfficialName(xlanguage.Arabic, "الجمهورية اليمنية")
	dataYemen.RegisterCapital(xlanguage.Arabic, "صنعاء")
}
