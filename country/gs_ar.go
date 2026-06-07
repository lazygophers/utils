//go:build (lang_ar || lang_all) && (country_all || country_antarctic || country_gs)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Arabic, "جورجيا الجنوبية وجزر ساندويتش الجنوبية")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Arabic, "إقليم جورجيا الجنوبية وجزر ساندويتش الجنوبية")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Arabic, "غريتفيكن")
}
