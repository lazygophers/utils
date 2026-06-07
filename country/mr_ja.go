//go:build (lang_ja || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Japanese, "モーリタニア")
	dataMauritania.RegisterOfficialName(xlanguage.Japanese, "モーリタニア・イスラム共和国")
	dataMauritania.RegisterCapital(xlanguage.Japanese, "ヌアクショット")
}
