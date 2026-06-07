//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Japanese, "イスラエル")
	dataIsrael.RegisterOfficialName(xlanguage.Japanese, "イスラエル国")
	dataIsrael.RegisterCapital(xlanguage.Japanese, "エルサレム")
}
