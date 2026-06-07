//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_pn || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Arabic, "جزر بيتكيرن")
	dataPitcairn.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر بيتكيرن وهندرسون ودوسي وأوينو")
	dataPitcairn.RegisterCapital(xlanguage.Arabic, "أدامزتاون")
}
