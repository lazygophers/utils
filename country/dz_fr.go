//go:build (lang_fr || lang_all) && (country_africa || country_all || country_dz || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.French, "Algérie")
	dataAlgeria.RegisterOfficialName(xlanguage.French, "République algérienne démocratique et populaire")
	dataAlgeria.RegisterCapital(xlanguage.French, "Alger")
}
