//go:build country_africa || country_all || country_northern_africa || country_sd

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Arabic, "السودان")
	dataSudan.RegisterOfficialName(xlanguage.Arabic, "جمهورية السودان")
	dataSudan.RegisterCapital(xlanguage.Arabic, "الخرطوم")
}
