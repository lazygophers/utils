package language

import (
	"net/http"
	"testing"

	xlanguage "golang.org/x/text/language"
)

func TestMake(t *testing.T) {
	tag := Make("zh-CN")
	if tag.String() != "zh-CN" {
		t.Errorf("got %q, want %q", tag.String(), "zh-CN")
	}
	if tag.Weight() != 0 {
		t.Errorf("default weight should be 0, got %f", tag.Weight())
	}
}

func TestParse(t *testing.T) {
	tag, err := Parse("en-US")
	if err != nil {
		t.Fatal(err)
	}
	if tag.Base() != "en" {
		t.Errorf("base: got %q, want %q", tag.Base(), "en")
	}
	if tag.Region() != "US" {
		t.Errorf("region: got %q, want %q", tag.Region(), "US")
	}
}

func TestTag_StandardConversion(t *testing.T) {
	tag := Make("zh-Hant-TW")
	xt := tag.Tag()

	supported := []xlanguage.Tag{xlanguage.English, xt}
	matcher := xlanguage.NewMatcher(supported)
	best, _, _ := matcher.Match(xlanguage.SimplifiedChinese)
	if best.String() != "zh-Hant-TW" {
		t.Errorf("matcher: got %q", best.String())
	}
}

func TestTag_BaseRegionScript(t *testing.T) {
	tag := Make("zh-Hant-TW")
	if tag.Base() != "zh" {
		t.Errorf("base: got %q", tag.Base())
	}
	if tag.Script() != "Hant" {
		t.Errorf("script: got %q", tag.Script())
	}
	if tag.Region() != "TW" {
		t.Errorf("region: got %q", tag.Region())
	}
}

func TestTag_Parent(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"zh-CN", "zh"},
		{"zh-TW", "zh-Hant"},
		{"en-US", "en"},
		{"en", "und"},
		{"und", "und"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			p := Make(tc.input).Parent()
			if p.String() != tc.want {
				t.Errorf("Parent(%s) = %q, want %q", tc.input, p.String(), tc.want)
			}
		})
	}
}

func TestTag_FallbackChain(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"zh-CN", []string{"zh-CN", "zh", "und"}},
		{"zh-TW", []string{"zh-TW", "zh-Hant", "und"}},
		{"en-US", []string{"en-US", "en", "und"}},
		{"en", []string{"en", "und"}},
		{"und", []string{"und"}},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			chain := Make(tc.input).FallbackChain()
			if len(chain) != len(tc.want) {
				t.Fatalf("got %d, want %d", len(chain), len(tc.want))
			}
			for i, tag := range chain {
				if tag.String() != tc.want[i] {
					t.Errorf("chain[%d]: got %q, want %q", i, tag.String(), tc.want[i])
				}
			}
		})
	}
}

func TestTag_Match(t *testing.T) {
	tests := []struct {
		name   string
		from   string
		target string
		want   bool
	}{
		{"zh-CN matches zh", "zh-CN", "zh", true},
		{"zh-CN matches und", "zh-CN", "und", true},
		{"zh-CN not match en", "zh-CN", "en", false},
		{"en-US matches en", "en-US", "en", true},
		{"en matches en", "en", "en", true},
		{"zh not match zh-CN", "zh", "zh-CN", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Make(tc.from).Match(Make(tc.target))
			if got != tc.want {
				t.Errorf("Match(%s, %s) = %v, want %v", tc.from, tc.target, got, tc.want)
			}
		})
	}
}

func TestTag_IsRTL(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"ar", true},
		{"he", true},
		{"fa", true},
		{"en", false},
		{"zh-CN", false},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := Make(tc.input).IsRTL()
			if got != tc.want {
				t.Errorf("IsRTL(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestParseAcceptLanguage(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   []string
	}{
		{"empty", "", nil},
		{"single", "en", []string{"en"}},
		{"multiple", "da, en-gb;q=0.8, en;q=0.7", []string{"da", "en-GB", "en"}},
		{"unordered q", "en;q=0.5, ja, zh-CN;q=0.8", []string{"ja", "zh-CN", "en"}},
		{"q=0 excluded", "en;q=0", nil},
		{"invalid skipped", "en, INVALID_TAG_XXX, zh", []string{"en", "zh"}},
		{"spaces", " en-US , zh-CN ; q=0.9 , ja ; q=0.8 ", []string{"en-US", "zh-CN", "ja"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ParseAcceptLanguage(tc.header)
			if len(got) != len(tc.want) {
				t.Fatalf("got %d tags, want %d: %+v", len(got), len(tc.want), got)
			}
			for i, tag := range got {
				if tag.String() != tc.want[i] {
					t.Errorf("tag[%d]: got %q, want %q", i, tag.String(), tc.want[i])
				}
			}
		})
	}
}

func TestParseAcceptLanguage_Weight(t *testing.T) {
	tags := ParseAcceptLanguage("ja;q=0.9, zh-CN;q=0.8, en;q=0.5")
	if len(tags) != 3 {
		t.Fatalf("got %d tags", len(tags))
	}
	if tags[0].Weight() != 0.9 {
		t.Errorf("tags[0] weight: got %f, want 0.9", tags[0].Weight())
	}
	if tags[1].Weight() != 0.8 {
		t.Errorf("tags[1] weight: got %f, want 0.8", tags[1].Weight())
	}
	if tags[2].Weight() != 0.5 {
		t.Errorf("tags[2] weight: got %f, want 0.5", tags[2].Weight())
	}
	tags2 := ParseAcceptLanguage("en")
	if tags2[0].Weight() != 1.0 {
		t.Errorf("default weight: got %f, want 1.0", tags2[0].Weight())
	}
}

func TestDetect(t *testing.T) {
	supported := []*Tag{
		Make("en"),
		Make("zh-CN"),
		Make("zh-TW"),
		Make("ja"),
	}

	tests := []struct {
		name      string
		header    string
		wantTag   string
		wantIndex int
	}{
		{"exact zh-CN", "zh-CN", "zh-CN", 1},
		{"fallback", "ko", "en", 0},
		{"quality order", "zh-TW;q=0.9, en;q=0.8", "zh-TW", 2},
		{"empty header", "", "en", 0},
		{"complex", "ja;q=0.9, zh-CN;q=0.8, en;q=0.5", "ja", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tag, idx := Detect(tc.header, supported)
			if idx != tc.wantIndex {
				t.Errorf("index: got %d, want %d", idx, tc.wantIndex)
			}
			if tag.String() != tc.wantTag {
				t.Errorf("tag: got %q, want %q", tag.String(), tc.wantTag)
			}
		})
	}
}

func TestDetect_EmptySupported(t *testing.T) {
	tag, idx := Detect("en", nil)
	if idx != -1 {
		t.Errorf("expected -1, got %d", idx)
	}
	if tag != nil {
		t.Errorf("expected nil, got %v", tag)
	}
}

func TestDetectFromStrings(t *testing.T) {
	tag, idx := DetectFromStrings("zh-CN, en;q=0.8", []string{"en", "zh-CN"})
	if idx != 1 {
		t.Errorf("index: got %d, want 1", idx)
	}
	if tag.Base() != "zh" {
		t.Errorf("base: got %q, want zh", tag.Base())
	}
}

func TestDetect_HTTPIntegration(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	supported := []*Tag{Make("en"), Make("zh-CN")}
	tag, idx := Detect(req.Header.Get("Accept-Language"), supported)

	if idx != 1 {
		t.Errorf("expected zh-CN (index 1), got index %d tag %q", idx, tag.String())
	}
	if !tag.Match(Make("zh")) {
		t.Errorf("expected Match(zh)=true")
	}
}
