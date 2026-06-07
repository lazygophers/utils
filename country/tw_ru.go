//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Russian, "Тайвань")
	dataTaiwan.RegisterOfficialName(xlanguage.Russian, "Китайская Республика (Тайвань)")
	dataTaiwan.RegisterCapital(xlanguage.Russian, "Тайбэй")
}
