//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Japanese, "カンボジア")
	dataCambodia.RegisterOfficialName(xlanguage.Japanese, "カンボジア王国")
	dataCambodia.RegisterCapital(xlanguage.Japanese, "プノンペン")
}
