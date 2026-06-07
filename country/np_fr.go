//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.French, "Népal")
	dataNepal.RegisterOfficialName(xlanguage.French, "République fédérale démocratique du Népal")
	dataNepal.RegisterCapital(xlanguage.French, "Katmandou")
}
