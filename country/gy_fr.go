//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.French, "Guyana")
	dataGuyana.RegisterOfficialName(xlanguage.French, "République coopérative du Guyana")
	dataGuyana.RegisterCapital(xlanguage.French, "Georgetown")
}
