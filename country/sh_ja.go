//go:build (lang_ja || lang_all) && (country_africa || country_all || country_sh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Japanese, "セントヘレナ")
	dataSaintHelena.RegisterOfficialName(xlanguage.Japanese, "セントヘレナ・アセンションおよびトリスタンダクーニャ")
	dataSaintHelena.RegisterCapital(xlanguage.Japanese, "ジェームズタウン")
}
