//go:build (lang_ar || lang_all) && (country_all || country_asia || country_az || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Arabic, "أذربيجان")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Arabic, "جمهورية أذربيجان")
	dataAzerbaijan.RegisterCapital(xlanguage.Arabic, "باكو")
}
