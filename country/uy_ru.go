//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Russian, "Уругвай")
	dataUruguay.RegisterOfficialName(xlanguage.Russian, "Восточная Республика Уругвай")
	dataUruguay.RegisterCapital(xlanguage.Russian, "Монтевидео")
}
