//go:build (lang_ja || lang_all) && (country_all || country_ba || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Japanese, "ボスニア・ヘルツェゴビナ")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Japanese, "ボスニア・ヘルツェゴビナ")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Japanese, "サラエヴォ")
}
