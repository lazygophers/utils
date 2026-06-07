//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Russian, "Парагвай")
	dataParaguay.RegisterOfficialName(xlanguage.Russian, "Республика Парагвай")
	dataParaguay.RegisterCapital(xlanguage.Russian, "Асунсьон")
}
