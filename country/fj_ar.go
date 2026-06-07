//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Arabic, "فيجي")
	dataFiji.RegisterOfficialName(xlanguage.Arabic, "جمهورية فيجي")
	dataFiji.RegisterCapital(xlanguage.Arabic, "سوفا")
}
