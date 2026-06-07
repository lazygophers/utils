//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Japanese, "エリトリア")
	dataEritrea.RegisterOfficialName(xlanguage.Japanese, "エリトリア国")
	dataEritrea.RegisterCapital(xlanguage.Japanese, "アスマラ")
}
