//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Arabic, "رومانيا")
	dataRomania.RegisterOfficialName(xlanguage.Arabic, "رومانيا")
	dataRomania.RegisterCapital(xlanguage.Arabic, "بوخارست")
}
