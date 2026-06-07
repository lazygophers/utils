//go:build (lang_ar || lang_all) && (country_all || country_americas || country_central_america || country_pa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Arabic, "بنما")
	dataPanama.RegisterOfficialName(xlanguage.Arabic, "جمهورية بنما")
	dataPanama.RegisterCapital(xlanguage.Arabic, "مدينة بنما")
}
