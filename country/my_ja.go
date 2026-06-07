//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Japanese, "マレーシア")
	dataMalaysia.RegisterOfficialName(xlanguage.Japanese, "マレーシア")
	dataMalaysia.RegisterCapital(xlanguage.Japanese, "クアラルンプール")
}
