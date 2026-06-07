//go:build (lang_es || lang_all) && (country_all || country_americas || country_bq || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Spanish, "Bonaire, San Eustaquio y Saba")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Spanish, "Bonaire, San Eustaquio y Saba")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.Spanish, "Kralendijk")
}
