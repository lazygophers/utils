//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Arabic, "جزر مارشال")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Arabic, "جمهورية جزر مارشال")
	dataMarshallIslands.RegisterCapital(xlanguage.Arabic, "ماجورو")
}
