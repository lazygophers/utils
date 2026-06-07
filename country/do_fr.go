//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.French, "République dominicaine")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.French, "République dominicaine")
	dataDominicanRepublic.RegisterCapital(xlanguage.French, "Saint-Domingue")
}
