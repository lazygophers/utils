//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.French, "Irlande")
	dataIreland.RegisterOfficialName(xlanguage.French, "Irlande")
	dataIreland.RegisterCapital(xlanguage.French, "Dublin")
}
