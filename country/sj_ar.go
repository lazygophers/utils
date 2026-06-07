//go:build (lang_ar || lang_all) && (country_all || country_europe || country_northern_europe || country_sj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Arabic, "سفالبارد ويان ماين")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Arabic, "سفالبارد ويان ماين")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Arabic, "لونغييربين")
}
