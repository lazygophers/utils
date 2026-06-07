//go:build country_all || country_asia || country_western_asia || country_ye

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Chinese, "也门")
	dataYemen.RegisterOfficialName(xlanguage.Chinese, "也门共和国")
	dataYemen.RegisterCapital(xlanguage.Chinese, "萨那")
}
