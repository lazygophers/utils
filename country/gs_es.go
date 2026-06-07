//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Spanish, "Islas Georgias del Sur y Sandwich del Sur")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Spanish, "Islas Georgias del Sur y Sandwich del Sur")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Spanish, "King Edward Point")
}
