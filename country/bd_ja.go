//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.Japanese, "バングラデシュ")
	dataBangladesh.RegisterOfficialName(xlanguage.Japanese, "バングラデシュ人民共和国")
	dataBangladesh.RegisterCapital(xlanguage.Japanese, "ダッカ")
}
