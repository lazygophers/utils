//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Japanese, "エルサルバドル")
	dataElSalvador.RegisterOfficialName(xlanguage.Japanese, "エルサルバドル共和国")
	dataElSalvador.RegisterCapital(xlanguage.Japanese, "サンサルバドル")
}
