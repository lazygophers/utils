//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.French, "Grèce")
	dataGreece.RegisterOfficialName(xlanguage.French, "République hellénique")
	dataGreece.RegisterCapital(xlanguage.French, "Athènes")
}
