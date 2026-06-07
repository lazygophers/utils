//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Japanese, "セントヘレナ")
	dataSaintHelena.RegisterOfficialName(xlanguage.Japanese, "セントヘレナ・アセンションおよびトリスタンダクーニャ")
	dataSaintHelena.RegisterCapital(xlanguage.Japanese, "ジェームズタウン")
}
