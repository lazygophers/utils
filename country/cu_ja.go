//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Japanese, "キューバ")
	dataCuba.RegisterOfficialName(xlanguage.Japanese, "キューバ共和国")
	dataCuba.RegisterCapital(xlanguage.Japanese, "ハバナ")
}
