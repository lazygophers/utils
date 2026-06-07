//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Japanese, "合衆国領有小離島")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Japanese, "合衆国領有小離島")
}
