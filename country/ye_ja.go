//go:build (lang_ja || lang_all) && (country_all || country_asia || country_western_asia || country_ye)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Japanese, "イエメン")
	dataYemen.RegisterOfficialName(xlanguage.Japanese, "イエメン共和国")
	dataYemen.RegisterCapital(xlanguage.Japanese, "サナア")
}
