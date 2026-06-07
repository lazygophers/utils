//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Arabic, "جمهورية أفريقيا الوسطى")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Arabic, "جمهورية أفريقيا الوسطى")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Arabic, "بانغي")
}
