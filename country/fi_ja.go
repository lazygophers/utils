//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Japanese, "フィンランド")
	dataFinland.RegisterOfficialName(xlanguage.Japanese, "フィンランド共和国")
	dataFinland.RegisterCapital(xlanguage.Japanese, "ヘルシンキ")
}
