//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.Arabic, "بوتسوانا")
	dataBotswana.RegisterOfficialName(xlanguage.Arabic, "جمهورية بوتسوانا")
	dataBotswana.RegisterCapital(xlanguage.Arabic, "غابورون")
}
