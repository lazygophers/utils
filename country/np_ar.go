//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Arabic, "نيبال")
	dataNepal.RegisterOfficialName(xlanguage.Arabic, "جمهورية نيبال الديمقراطية الاتحادية")
	dataNepal.RegisterCapital(xlanguage.Arabic, "كاتماندو")
}
