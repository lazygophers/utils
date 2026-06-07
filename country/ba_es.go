//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Spanish, "Bosnia y Herzegovina")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Spanish, "Bosnia y Herzegovina")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Spanish, "Sarajevo")
}
