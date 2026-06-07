//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Japanese, "スペイン")
	dataSpain.RegisterOfficialName(xlanguage.Japanese, "スペイン王国")
	dataSpain.RegisterCapital(xlanguage.Japanese, "マドリード")
}
