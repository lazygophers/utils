//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Japanese, "台湾")
	dataTaiwan.RegisterOfficialName(xlanguage.Japanese, "中華民国")
	dataTaiwan.RegisterCapital(xlanguage.Japanese, "台北")
}
