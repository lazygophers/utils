//go:build country_all || country_americas || country_bq || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.English, "Caribbean Netherlands")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.English, "Bonaire, Sint Eustatius and Saba")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.English, "Kralendijk")
}
