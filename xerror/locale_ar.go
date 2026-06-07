//go:build lang_ar || lang_all

package xerror

// 阿拉伯文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("ar", map[int]string{
		CodeSystem:        "خطأ في النظام",
		CodeInvalidParam:  "معامل غير صالح",
		CodeNoAuth:        "غير مصرح به",
		CodeNoData:        "غير موجود",
		CodeConflict:      "تعارض في البيانات",
		CodeNotLogin:      "لم يتم تسجيل الدخول",
		CodeTimeout:       "انتهت المهلة",
		CodeRateLimited:   "طلبات كثيرة جدًا",
		CodeForbidden:     "ممنوع",
		CodeUnavailable:   "الخدمة غير متاحة",
		CodeDataCorrupted: "البيانات تالفة",
	})
}
