//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Japanese, "コモロ")
	dataComoros.RegisterOfficialName(xlanguage.Japanese, "コモロ連合")
	dataComoros.RegisterCapital(xlanguage.Japanese, "モロニ")
}
