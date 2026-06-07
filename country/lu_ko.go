//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Korean, "룩셈부르크")
	dataLuxembourg.RegisterOfficialName(xlanguage.Korean, "룩셈부르크 대공국")
	dataLuxembourg.RegisterCapital(xlanguage.Korean, "룩셈부르크")
}
