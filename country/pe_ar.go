//go:build (lang_ar || lang_all) && (country_all || country_americas || country_pe || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Arabic, "بيرو")
	dataPeru.RegisterOfficialName(xlanguage.Arabic, "جمهورية بيرو")
	dataPeru.RegisterCapital(xlanguage.Arabic, "ليما")
}
