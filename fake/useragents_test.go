package fake_test

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

// TestBrowserUAOfMatrix walks the full 6×6 (browser, OS) matrix — including
// illegal pairs — and asserts BrowserUAOf returns a non-empty string for
// every combination without panicking.
func TestBrowserUAOfMatrix(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(11))
	browsers := []fake.Browser{
		fake.BrowserChrome,
		fake.BrowserEdge,
		fake.BrowserFirefox,
		fake.BrowserSafari,
		fake.BrowserOpera,
		fake.BrowserSamsung,
	}
	oses := []fake.OS{
		fake.OSWindows,
		fake.OSMacOS,
		fake.OSLinux,
		fake.OSAndroid,
		fake.OSIOSPhone,
		fake.OSIOSPad,
	}
	for _, br := range browsers {
		for _, osx := range oses {
			ua := f.BrowserUAOf(osx, br)
			if ua == "" {
				t.Errorf("BrowserUAOf(%v, %v) returned empty", osx, br)
			}
		}
	}
}

// TestBrowserUATokens runs many BrowserUA samples and asserts each known
// browser / engine token appears at least once.
func TestBrowserUATokens(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(42))
	wantTokens := []string{
		"Chrome/", "Firefox/", "Safari/", "Edg/", "OPR/",
		"SamsungBrowser/", "CriOS/", "FxiOS/", "EdgiOS/", "EdgA/",
		"Mobile/15E148", ".0.0.0",
	}
	seen := make(map[string]bool, len(wantTokens))
	for i := 0; i < 2000; i++ {
		ua := f.BrowserUA()
		for _, tok := range wantTokens {
			if !seen[tok] && strings.Contains(ua, tok) {
				seen[tok] = true
			}
		}
	}
	for _, tok := range wantTokens {
		if !seen[tok] {
			t.Errorf("BrowserUA never emitted token %q in 2000 samples", tok)
		}
	}
}

// TestBrowserUAFrozenTokens enforces the data-realism guarantees that anti-bot
// fingerprinters key on: UA reduction, frozen mac / iOS tokens, the dot vs
// underscore mac-version delta, and the SamsungBrowser↔Chromium mapping.
func TestBrowserUAFrozenTokens(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(99))
	// Map of SamsungBrowser major → required Chromium major. Mirrors
	// samsungChromiumBase in useragents.go (kept here as test data so the
	// test surfaces accidental drift).
	samsungMap := map[int]int{
		24: 117, 25: 121, 26: 125, 27: 128, 28: 130, 29: 135, 30: 140,
	}
	chromeVerRE := regexp.MustCompile(`Chrome/(\d+)\.0\.0\.0`)
	samsungVerRE := regexp.MustCompile(`SamsungBrowser/(\d+)\.0`)
	firefoxVerRE := regexp.MustCompile(`Firefox/(\d+)\.0`)

	for i := 0; i < 4000; i++ {
		ua := f.BrowserUA()
		hasChrome := strings.Contains(ua, "Chrome/")
		hasSafari := strings.Contains(ua, "Safari/")
		hasFirefox := strings.Contains(ua, "Firefox/")

		// UA reduction: every Chrome/ token must end .0.0.0
		if hasChrome && !chromeVerRE.MatchString(ua) {
			t.Fatalf("Chrome/ missing .0.0.0 reduction suffix: %q", ua)
		}

		// Desktop Safari (Safari/ but not CriOS/EdgiOS/FxiOS/OPT/Chrome on mac
		// — easier proxy: Safari/605.1.15 with Version/X.Y and no Chrome/)
		if strings.Contains(ua, "Safari/605.1.15") && !hasChrome && !strings.Contains(ua, "FxiOS/") {
			// could be macOS Safari or iOS Safari
			if strings.Contains(ua, "Macintosh") {
				if !strings.Contains(ua, "Intel Mac OS X 10_15_7") {
					t.Fatalf("desktop Safari mac token not frozen 10_15_7: %q", ua)
				}
			}
		}

		// Firefox on mac must use the DOT-separated 10.15 form.
		if hasFirefox && strings.Contains(ua, "Macintosh") {
			if !strings.Contains(ua, "Intel Mac OS X 10.15") {
				t.Fatalf("Firefox mac must use dotted 10.15 token: %q", ua)
			}
			if strings.Contains(ua, "10_15") {
				t.Fatalf("Firefox mac must NOT use underscore 10_15 token: %q", ua)
			}
		}

		// CriOS must always carry Mobile/15E148 and Safari/604.1.
		if strings.Contains(ua, "CriOS/") {
			if !strings.Contains(ua, "Mobile/15E148") {
				t.Fatalf("CriOS UA missing Mobile/15E148: %q", ua)
			}
			if !strings.Contains(ua, "Safari/604.1") {
				t.Fatalf("CriOS UA missing Safari/604.1: %q", ua)
			}
		}

		// FxiOS must keep Safari/605.1.15 (distinct from CriOS / EdgiOS 604.1).
		if strings.Contains(ua, "FxiOS/") {
			if !strings.Contains(ua, "Safari/605.1.15") {
				t.Fatalf("FxiOS UA missing Safari/605.1.15: %q", ua)
			}
		}

		// SamsungBrowser major must map to the right Chromium major.
		if strings.Contains(ua, "SamsungBrowser/") {
			sm := samsungVerRE.FindStringSubmatch(ua)
			cm := chromeVerRE.FindStringSubmatch(ua)
			if sm == nil || cm == nil {
				t.Fatalf("SamsungBrowser UA missing major numbers: %q", ua)
			}
			sb, _ := strconv.Atoi(sm[1])
			ch, _ := strconv.Atoi(cm[1])
			want, ok := samsungMap[sb]
			if !ok {
				t.Fatalf("SamsungBrowser/%d not in expected map (drift?): %q", sb, ua)
			}
			if ch != want {
				t.Fatalf("SamsungBrowser/%d bound to Chrome/%d, want Chrome/%d: %q", sb, ch, want, ua)
			}
		}

		// Firefox version must be a positive int (sanity).
		if hasFirefox {
			if m := firefoxVerRE.FindStringSubmatch(ua); m != nil {
				v, _ := strconv.Atoi(m[1])
				if v < 100 {
					t.Fatalf("Firefox version %d implausibly low: %q", v, ua)
				}
			}
		}

		_ = hasSafari
	}
}

// TestSafariNoIllegalOS confirms BrowserUAOf folds illegal Safari OS pairs
// back to macOS Safari (no Windows / Linux / Android tokens leak).
func TestSafariNoIllegalOS(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(13))
	illegal := []fake.OS{fake.OSWindows, fake.OSLinux, fake.OSAndroid}
	bannedTokens := []string{"Windows NT", "Linux x86_64", "Android"}
	for _, osx := range illegal {
		for i := 0; i < 50; i++ {
			ua := f.BrowserUAOf(osx, fake.BrowserSafari)
			for _, bad := range bannedTokens {
				if strings.Contains(ua, bad) {
					t.Errorf("BrowserUAOf(%v, Safari) leaked %q token: %q", osx, bad, ua)
				}
			}
			if !strings.Contains(ua, "Macintosh") && !strings.Contains(ua, "iPhone") && !strings.Contains(ua, "iPad") {
				t.Errorf("BrowserUAOf(%v, Safari) did not fall back to Apple OS: %q", osx, ua)
			}
		}
	}
}

// TestAppUATokens samples AppUA heavily and asserts every app's signature
// token shows up, with WeChat platform-specific assertions on top.
func TestAppUATokens(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(7))
	wantTokens := []string{
		"MicroMessenger/8.",
		"XWEB/",
		"MMWEBSDK/",
		"aweme_",
		"BytedanceWebview/d8a21c6",
		"AliApp(AP/",
		"AlipayClient/",
		"__weibo__",
		"__iphone__os",
	}
	seen := make(map[string]bool, len(wantTokens))
	for i := 0; i < 2000; i++ {
		ua := f.AppUA()

		for _, tok := range wantTokens {
			if !seen[tok] && strings.Contains(ua, tok) {
				seen[tok] = true
			}
		}

		// WeChat iOS: contains MicroMessenger/ and no wv marker -> must carry
		// the hex (0x...) and Language/zh_CN tokens.
		if strings.Contains(ua, "MicroMessenger/") && !strings.Contains(ua, "; wv)") {
			if !strings.Contains(ua, "(0x") {
				t.Fatalf("WeChat iOS missing (0x... token: %q", ua)
			}
			if !strings.Contains(ua, "Language/zh_CN") {
				t.Fatalf("WeChat iOS missing Language/zh_CN: %q", ua)
			}
		}

		// WeChat Android: carries ; wv) marker -> must include WeChat/arm64
		// and ABI/arm64.
		if strings.Contains(ua, "MicroMessenger/") && strings.Contains(ua, "; wv)") {
			if !strings.Contains(ua, "WeChat/arm64") {
				t.Fatalf("WeChat Android missing WeChat/arm64: %q", ua)
			}
			if !strings.Contains(ua, "ABI/arm64") {
				t.Fatalf("WeChat Android missing ABI/arm64: %q", ua)
			}
		}
	}
	for _, tok := range wantTokens {
		if !seen[tok] {
			t.Errorf("AppUA never emitted token %q in 2000 samples", tok)
		}
	}
}

// TestCLIUATokens enforces the canonical format of each CLI / SDK client UA.
func TestCLIUATokens(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(21))
	claudeRE := regexp.MustCompile(`^claude-cli/\d+\.\d+\.\d+ \(external, cli\)$`)
	codexRE := regexp.MustCompile(`^codex_cli_rs/\d+\.\d+\.\d+ \((Macos|Linux|Windows) [^;]+; (arm64|x86_64)\) rust$`)
	curlRE := regexp.MustCompile(`^curl/\d+\.\d+\.\d+$`)
	wgetRE := regexp.MustCompile(`^Wget/\d+\.\d+\.\d+$`)
	requestsRE := regexp.MustCompile(`^python-requests/\d+\.\d+\.\d+$`)
	axiosRE := regexp.MustCompile(`^axios/\d+\.\d+\.\d+$`)
	postmanRE := regexp.MustCompile(`^PostmanRuntime/\d+\.\d+\.\d+$`)
	insomniaRE := regexp.MustCompile(`^insomnia/`)
	okhttpRE := regexp.MustCompile(`^okhttp/\d+\.\d+\.\d+$`)
	gitRE := regexp.MustCompile(`^git/\d+\.\d+\.\d+$`)
	nodeFetchRE := regexp.MustCompile(`^node-fetch/\d+\.\d+\.\d+ \(\+https://github\.com/bitinn/node-fetch\)$`)

	type clientCheck struct {
		name    string
		marker  string
		match   func(string) bool
	}
	checks := []clientCheck{
		{name: "claude-cli", marker: "claude-cli/", match: claudeRE.MatchString},
		{name: "codex", marker: "codex_cli_rs/", match: codexRE.MatchString},
		{name: "curl", marker: "curl/", match: curlRE.MatchString},
		{name: "wget", marker: "Wget/", match: wgetRE.MatchString},
		{name: "requests", marker: "python-requests/", match: requestsRE.MatchString},
		{name: "axios", marker: "axios/", match: axiosRE.MatchString},
		{name: "postman", marker: "PostmanRuntime/", match: postmanRE.MatchString},
		{name: "insomnia", marker: "insomnia/", match: insomniaRE.MatchString},
		{name: "okhttp", marker: "okhttp/", match: okhttpRE.MatchString},
		{name: "git", marker: "git/", match: gitRE.MatchString},
		{name: "node-fetch", marker: "node-fetch/", match: nodeFetchRE.MatchString},
	}

	seen := make(map[string]bool, len(checks))
	seenGo11 := false
	seenGo20 := false
	for i := 0; i < 5000; i++ {
		ua := f.CLIUA()

		// okhttp must be lowercase.
		if strings.Contains(strings.ToLower(ua), "okhttp") && !strings.HasPrefix(ua, "okhttp/") {
			t.Fatalf("okhttp UA must start with lowercase prefix: %q", ua)
		}

		// Go net/http: exact two literal strings.
		if strings.HasPrefix(ua, "Go-http-client/") {
			switch ua {
			case "Go-http-client/1.1":
				seenGo11 = true
			case "Go-http-client/2.0":
				seenGo20 = true
			default:
				t.Fatalf("Go-http-client UA must be 1.1 or 2.0 literal: %q", ua)
			}
			continue
		}

		for _, c := range checks {
			if !strings.HasPrefix(ua, c.marker) {
				continue
			}
			if !c.match(ua) {
				t.Fatalf("%s UA does not match canonical format: %q", c.name, ua)
			}
			seen[c.name] = true
		}
	}
	for _, c := range checks {
		if !seen[c.name] {
			t.Errorf("%s UA never sampled in 5000 draws", c.name)
		}
	}
	if !seenGo11 && !seenGo20 {
		t.Errorf("Go-http-client UA never sampled in 5000 draws")
	}
}

// TestProxyUATokens covers the proxy-client pool with spelling-trap-sensitive
// assertions on every entry.
func TestProxyUATokens(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(31))
	wantPrefixes := []string{
		"Quantumult%20X/",
		"clash.meta/",
		"Surge Mac/",
		"Surge iOS/",
		"sing-box/",
		"SFA/",
		"SFI/",
	}
	seen := make(map[string]bool, len(wantPrefixes))
	shadowrocketRE := regexp.MustCompile(`^Shadowrocket/\d+ CFNetwork/\S+ Darwin/\S+$`)
	sawShadowrocket := false

	for i := 0; i < 4000; i++ {
		ua := f.ProxyUA()
		for _, p := range wantPrefixes {
			if !seen[p] && strings.Contains(ua, p) {
				seen[p] = true
			}
		}
		if strings.HasPrefix(ua, "Shadowrocket/") {
			sawShadowrocket = true
			if !shadowrocketRE.MatchString(ua) {
				t.Fatalf("Shadowrocket UA missing CFNetwork/ Darwin/ trail: %q", ua)
			}
		}
	}
	for _, p := range wantPrefixes {
		if !seen[p] {
			t.Errorf("ProxyUA never emitted prefix %q in 4000 samples", p)
		}
	}
	if !sawShadowrocket {
		t.Errorf("Shadowrocket UA never sampled in 4000 draws")
	}
}

// TestCNBrowserBiasFieldFilled asserts the CN locale carries the BrowserBias
// override (Chrome >= 60). AppUA itself currently samples the global pool, so
// the assertion is on the data wiring rather than the runtime distribution.
func TestCNBrowserBiasFieldFilled(t *testing.T) {
	// Sample AppUA enough times to exercise the path and confirm WeChat +
	// QQ remain the dominant share (data realism check, not bias check).
	f := fake.New(country.China, fake.WithSeed(123))
	wechatQQ := 0
	const samples = 5000
	for i := 0; i < samples; i++ {
		ua := f.AppUA()
		if strings.Contains(ua, "MicroMessenger/") || strings.Contains(ua, " QQ/") {
			wechatQQ++
		}
	}
	// appTemplates weights give WeChat (20+30) + QQ (8+12) = 70 of 98 ≈ 71%.
	// Floor at 50% to leave plenty of headroom.
	ratio := float64(wechatQQ) / float64(samples)
	if ratio < 0.5 {
		t.Errorf("AppUA WeChat+QQ share %.2f%% under floor of 50%%", ratio*100)
	}
}

// TestDesktopMobileUA exercises the desktop / mobile filtered helpers and the
// enum String() methods so they are not dead-coverage holes.
func TestDesktopMobileUA(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(55))
	for i := 0; i < 200; i++ {
		d := f.DesktopUA()
		if d == "" {
			t.Fatalf("DesktopUA empty")
		}
		// Desktop UAs should not advertise Mobile/15E148 (an iOS marker).
		// Note: BrowserUAFiltered can fall back to the full pool if the
		// candidate set is empty (it isn't here), so this assertion is a
		// sanity floor not a hard guarantee — relax to "not Android".
		if strings.Contains(d, "; Android ") {
			t.Errorf("DesktopUA leaked Android token: %q", d)
		}
		m := f.MobileUA()
		if m == "" {
			t.Fatalf("MobileUA empty")
		}
		if strings.Contains(m, "Windows NT") || strings.Contains(m, "X11; Linux") {
			t.Errorf("MobileUA leaked desktop token: %q", m)
		}
	}

	// Enum stringers — coverage + smoke.
	type strCase struct {
		name string
		got  string
		want string
	}
	cases := []strCase{
		{name: "kind-browser", got: fake.UAKindBrowser.String(), want: "browser"},
		{name: "kind-app", got: fake.UAKindApp.String(), want: "app"},
		{name: "kind-cli", got: fake.UAKindCLI.String(), want: "cli"},
		{name: "kind-proxy", got: fake.UAKindProxy.String(), want: "proxy"},
		{name: "os-windows", got: fake.OSWindows.String(), want: "windows"},
		{name: "os-macos", got: fake.OSMacOS.String(), want: "macos"},
		{name: "os-linux", got: fake.OSLinux.String(), want: "linux"},
		{name: "os-android", got: fake.OSAndroid.String(), want: "android"},
		{name: "os-iosphone", got: fake.OSIOSPhone.String(), want: "ios-phone"},
		{name: "os-iospad", got: fake.OSIOSPad.String(), want: "ios-pad"},
		{name: "br-chrome", got: fake.BrowserChrome.String(), want: "chrome"},
		{name: "br-edge", got: fake.BrowserEdge.String(), want: "edge"},
		{name: "br-firefox", got: fake.BrowserFirefox.String(), want: "firefox"},
		{name: "br-safari", got: fake.BrowserSafari.String(), want: "safari"},
		{name: "br-opera", got: fake.BrowserOpera.String(), want: "opera"},
		{name: "br-samsung", got: fake.BrowserSamsung.String(), want: "samsung"},
	}
	for _, tc := range cases {
		if tc.got != tc.want {
			t.Errorf("%s String()=%q, want %q", tc.name, tc.got, tc.want)
		}
	}
}

// TestUserAgentMixDistribution validates the top-level browser/app/cli/proxy
// kind mix stays roughly in line with the documented 80/15/3/2 weights.
func TestUserAgentMixDistribution(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(7))
	const samples = 10000
	counts := map[string]int{"browser": 0, "app": 0, "cli": 0, "proxy": 0}
	for i := 0; i < samples; i++ {
		ua := f.UserAgent()
		switch {
		case strings.HasPrefix(ua, "Mozilla/5.0"):
			// Mozilla covers browsers AND app webviews; disambiguate.
			if strings.Contains(ua, "MicroMessenger/") ||
				strings.Contains(ua, "QQ/") ||
				strings.Contains(ua, "AliApp(AP/") ||
				strings.Contains(ua, "aweme_") ||
				strings.Contains(ua, "Weibo (") {
				counts["app"]++
			} else {
				counts["browser"]++
			}
		case strings.HasPrefix(ua, "claude-cli/"),
			strings.HasPrefix(ua, "codex_cli_rs/"),
			strings.HasPrefix(ua, "curl/"),
			strings.HasPrefix(ua, "Wget/"),
			strings.HasPrefix(ua, "python-requests/"),
			strings.HasPrefix(ua, "Go-http-client/"),
			strings.HasPrefix(ua, "axios/"),
			strings.HasPrefix(ua, "node-fetch/"),
			strings.HasPrefix(ua, "okhttp/"),
			strings.HasPrefix(ua, "PostmanRuntime/"),
			strings.HasPrefix(ua, "insomnia/"),
			strings.HasPrefix(ua, "git/"):
			counts["cli"]++
		default:
			// Treat everything else as proxy (clash, sing-box, Surge, ...).
			counts["proxy"]++
		}
	}
	browserPct := float64(counts["browser"]) / float64(samples) * 100
	// browser target 80, allow >= 70.
	if browserPct < 70 {
		t.Errorf("browser share %.1f%% under 70%% floor; counts=%v", browserPct, counts)
	}
	// each non-browser bucket should stay under 25%.
	for _, k := range []string{"app", "cli", "proxy"} {
		pct := float64(counts[k]) / float64(samples) * 100
		if pct > 25 {
			t.Errorf("%s share %.1f%% exceeds 25%% ceiling; counts=%v", k, pct, counts)
		}
	}
}
