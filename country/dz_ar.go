//go:build country_africa || country_all || country_dz || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Arabic, "الجزائر")
	dataAlgeria.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الجزائرية الديمقراطية الشعبية")
	dataAlgeria.RegisterCapital(xlanguage.Arabic, "الجزائر")
}
