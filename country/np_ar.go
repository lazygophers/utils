//go:build (lang_ar || lang_all) && (country_all || country_asia || country_np || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Arabic, "نيبال")
	dataNepal.RegisterOfficialName(xlanguage.Arabic, "جمهورية نيبال الديمقراطية الاتحادية")
	dataNepal.RegisterCapital(xlanguage.Arabic, "كاتماندو")
}
