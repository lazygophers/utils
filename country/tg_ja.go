//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Japanese, "トーゴ")
	dataTogo.RegisterOfficialName(xlanguage.Japanese, "トーゴ共和国")
	dataTogo.RegisterCapital(xlanguage.Japanese, "ロメ")
}
