//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Japanese, "キプロス")
	dataCyprus.RegisterOfficialName(xlanguage.Japanese, "キプロス共和国")
	dataCyprus.RegisterCapital(xlanguage.Japanese, "ニコシア")
}
