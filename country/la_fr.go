//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.French, "Laos")
	dataLaos.RegisterOfficialName(xlanguage.French, "République démocratique populaire lao")
	dataLaos.RegisterCapital(xlanguage.French, "Vientiane")
}
