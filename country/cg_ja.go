//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Japanese, "コンゴ共和国")
	dataCongo.RegisterOfficialName(xlanguage.Japanese, "コンゴ共和国")
	dataCongo.RegisterCapital(xlanguage.Japanese, "ブラザヴィル")
}
