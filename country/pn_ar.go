//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Arabic, "جزر بيتكيرن")
	dataPitcairn.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر بيتكيرن وهندرسون ودوسي وأوينو")
	dataPitcairn.RegisterCapital(xlanguage.Arabic, "أدامزتاون")
}
