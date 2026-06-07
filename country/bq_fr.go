//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.French, "Pays-Bas caribéens")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.French, "Bonaire, Saint-Eustache et Saba")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.French, "Kralendijk")
}
