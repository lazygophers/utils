//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Spanish, "Sint Maarten")
	dataSintMaarten.RegisterOfficialName(xlanguage.Spanish, "Sint Maarten")
	dataSintMaarten.RegisterCapital(xlanguage.Spanish, "Philipsburg")
}
