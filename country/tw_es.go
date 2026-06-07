//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Spanish, "Taiwán")
	dataTaiwan.RegisterOfficialName(xlanguage.Spanish, "República de China (Taiwán)")
	dataTaiwan.RegisterCapital(xlanguage.Spanish, "Taipéi")
}
