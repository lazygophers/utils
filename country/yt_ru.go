//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Russian, "Майотта")
	dataMayotte.RegisterOfficialName(xlanguage.Russian, "Департамент Майотта")
	dataMayotte.RegisterCapital(xlanguage.Russian, "Мамудзу")
}
