//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Russian, "Джибути")
	dataDjibouti.RegisterOfficialName(xlanguage.Russian, "Республика Джибути")
	dataDjibouti.RegisterCapital(xlanguage.Russian, "Джибути")
}
