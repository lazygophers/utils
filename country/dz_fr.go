//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.French, "Algérie")
	dataAlgeria.RegisterOfficialName(xlanguage.French, "République algérienne démocratique et populaire")
	dataAlgeria.RegisterCapital(xlanguage.French, "Alger")
}
