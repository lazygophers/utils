//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Arabic, "ليتوانيا")
	dataLithuania.RegisterOfficialName(xlanguage.Arabic, "جمهورية ليتوانيا")
	dataLithuania.RegisterCapital(xlanguage.Arabic, "فيلنيوس")
}
