//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Arabic, "موناكو")
	dataMonaco.RegisterOfficialName(xlanguage.Arabic, "إمارة موناكو")
	dataMonaco.RegisterCapital(xlanguage.Arabic, "موناكو")
}
