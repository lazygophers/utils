//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bq || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.French, "Pays-Bas caribéens")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.French, "Bonaire, Saint-Eustache et Saba")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.French, "Kralendijk")
}
