//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.French, "Géorgie du Sud-et-les îles Sandwich du Sud")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.French, "Géorgie du Sud-et-les îles Sandwich du Sud")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.French, "King Edward Point")
}
