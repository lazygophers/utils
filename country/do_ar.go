//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Arabic, "جمهورية الدومينيكان")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Arabic, "جمهورية الدومينيكان")
	dataDominicanRepublic.RegisterCapital(xlanguage.Arabic, "سانتو دومينغو")
}
