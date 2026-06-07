//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.Japanese, "スリランカ")
	dataSriLanka.RegisterOfficialName(xlanguage.Japanese, "スリランカ民主社会主義共和国")
	dataSriLanka.RegisterCapital(xlanguage.Japanese, "スリジャヤワルダナプラコッテ")
}
