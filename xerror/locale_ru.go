//go:build lang_ru || lang_all

package xerror

// 俄文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("ru", map[int]string{
		CodeSystem:        "системная ошибка",
		CodeInvalidParam:  "недопустимый параметр",
		CodeNoAuth:        "не авторизовано",
		CodeNoData:        "не найдено",
		CodeConflict:      "конфликт",
		CodeNotLogin:      "не выполнен вход",
		CodeTimeout:       "тайм-аут",
		CodeRateLimited:   "слишком много запросов",
		CodeForbidden:     "запрещено",
		CodeUnavailable:   "сервис недоступен",
		CodeDataCorrupted: "данные повреждены",
	})
}
