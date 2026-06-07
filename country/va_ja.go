//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Japanese, "バチカン市国")
	dataVaticanCity.RegisterOfficialName(xlanguage.Japanese, "バチカン市国")
	dataVaticanCity.RegisterCapital(xlanguage.Japanese, "バチカン")
}
