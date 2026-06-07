//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Japanese, "シンガポール")
	dataSingapore.RegisterOfficialName(xlanguage.Japanese, "シンガポール共和国")
	dataSingapore.RegisterCapital(xlanguage.Japanese, "シンガポール")
}
