//go:build (lang_ja || lang_all) && (country_all || country_asia || country_il || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Japanese, "イスラエル")
	dataIsrael.RegisterOfficialName(xlanguage.Japanese, "イスラエル国")
	dataIsrael.RegisterCapital(xlanguage.Japanese, "エルサレム")
}
