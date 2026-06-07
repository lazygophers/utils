//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.MustParse("zh-Hant"), "聖克里斯多福及尼維斯")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖克里斯多福及尼維斯聯邦")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴士地")
}
