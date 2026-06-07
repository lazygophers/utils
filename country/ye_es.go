//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Spanish, "Yemen")
	dataYemen.RegisterOfficialName(xlanguage.Spanish, "República de Yemen")
	dataYemen.RegisterCapital(xlanguage.Spanish, "Saná")
}
