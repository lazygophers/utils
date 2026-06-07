//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Spanish, "Zimbabue")
	dataZimbabwe.RegisterOfficialName(xlanguage.Spanish, "República de Zimbabue")
	dataZimbabwe.RegisterCapital(xlanguage.Spanish, "Harare")
}
