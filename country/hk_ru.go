//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Russian, "Гонконг")
	dataHongKong.RegisterOfficialName(xlanguage.Russian, "Специальный административный район Сянган Китайской Народной Республики")
	dataHongKong.RegisterCapital(xlanguage.Russian, "Гонконг")
}
