//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_kn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Arabic, "سانت كيتس ونيفيس")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Arabic, "اتحاد سانت كيتس ونيفيس")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Arabic, "باستير")
}
