//go:build (lang_ru || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Russian, "Мавритания")
	dataMauritania.RegisterOfficialName(xlanguage.Russian, "Исламская Республика Мавритания")
	dataMauritania.RegisterCapital(xlanguage.Russian, "Нуакшот")
}
