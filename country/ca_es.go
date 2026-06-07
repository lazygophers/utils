//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Spanish, "Canadá")
	dataCanada.RegisterOfficialName(xlanguage.Spanish, "Canadá")
	dataCanada.RegisterCapital(xlanguage.Spanish, "Ottawa")
}
