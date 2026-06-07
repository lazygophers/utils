//go:build lang_ko || lang_all

package xerror

// 韩文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("ko", map[int]string{
		CodeSystem:        "시스템 오류",
		CodeInvalidParam:  "잘못된 매개변수",
		CodeNoAuth:        "인증되지 않음",
		CodeNoData:        "데이터를 찾을 수 없음",
		CodeConflict:      "데이터 충돌",
		CodeNotLogin:      "로그인하지 않음",
		CodeTimeout:       "시간 초과",
		CodeRateLimited:   "요청이 너무 많음",
		CodeForbidden:     "접근 금지",
		CodeUnavailable:   "서비스 이용 불가",
		CodeDataCorrupted: "데이터 손상",
	})
}
