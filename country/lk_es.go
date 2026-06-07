//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.Spanish, "Sri Lanka")
	dataSriLanka.RegisterOfficialName(xlanguage.Spanish, "República Socialista Democrática de Sri Lanka")
	dataSriLanka.RegisterCapital(xlanguage.Spanish, "Sri Jayawardenapura Kotte")
}
