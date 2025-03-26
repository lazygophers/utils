package randx_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lazygophers/utils/anyx"
	"github.com/lazygophers/utils/candy"
	"github.com/lazygophers/utils/stringx"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

type UA struct {
	Useragent                string  `json:"useragent"`
	Percent                  float64 `json:"percent"`
	Type                     string  `json:"type"`
	DeviceBrand              string  `json:"device_brand"`
	Browser                  string  `json:"browser"`
	BrowserVersion           string  `json:"browser_version"`
	BrowserVersionMajorMinor float64 `json:"browser_version_major_minor"`
	Os                       string  `json:"os"`
	OsVersion                string  `json:"os_version"`
	Platform                 string  `json:"platform"`
}

func TestGenerate(t *testing.T) {
	//open := func() (io.Reader, error) {
	//	file, err := os.Open("browsers.jsonl")
	//	if err != nil {
	//		t.Errorf("err:%v", err)
	//		return nil, err
	//	}
	//
	//	return file, nil
	//}

	open := func() (io.Reader, error) {
		resp, err := http.Get("https://raw.githubusercontent.com/fake-useragent/fake-useragent/refs/heads/main/src/fake_useragent/data/browsers.jsonl")
		if err != nil {
			t.Errorf("err:%v", err)
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("status %s", resp.Status)
		}

		return resp.Body, nil
	}

	render, err := open()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	scanner := bufio.NewScanner(render)

	var uaList []*UA
	m := map[string][]*UA{}

	for scanner.Scan() {
		var ua UA
		err = json.Unmarshal(scanner.Bytes(), &ua)
		if err != nil {
			t.Errorf("err:%v", err)
			return
		}

		cover := func(str string) string {
			s := strings.ToLower(str)

			switch s {
			case "win32":
				s = "windows"
			case "generic_android_tablet", "generic_android":
				s = "android"
			}

			if strings.Contains(s, " ") {
				s = strings.ReplaceAll(s, "ui/wkwebview", "")
				s = strings.ReplaceAll(s, "ios", "")
				s = strings.ReplaceAll(s, "x86_64", "")
				s = strings.ReplaceAll(s, "armv81", "")
				s = strings.ReplaceAll(s, "aarch64", "")
				s = strings.ReplaceAll(s, "armv8l", "")
				s = strings.ReplaceAll(s, "mobile", "")
				s = strings.ReplaceAll(s, "browser", "")
				s = strings.ReplaceAll(s, "webview", "")
				s = strings.ReplaceAll(s, "internet", "")
			}

			return stringx.ToCamel(s)
		}

		ua.Type = cover(ua.Type)
		ua.Os = cover(ua.Os)
		ua.Platform = cover(ua.Platform)
		ua.DeviceBrand = cover(ua.DeviceBrand)
		ua.Browser = cover(ua.Browser)

		if ua.Os != "" {
			m[ua.Os] = append(m[ua.Os], &ua)
		}

		if ua.Type != "" {
			m[ua.Type] = append(m[ua.Type], &ua)
		}

		if ua.Browser != "" {
			m[ua.Browser] = append(m[ua.Browser], &ua)
		}

		if ua.Platform != "" {
			m[ua.Platform] = append(m[ua.Platform], &ua)
		}

		if ua.DeviceBrand != "" {
			m[ua.DeviceBrand] = append(m[ua.DeviceBrand], &ua)
		}

		uaList = append(uaList, &ua)
	}

	t.Log(anyx.MapKeysString(m))

	var b bytes.Buffer

	b.WriteString("package randx\n")
	b.WriteString("\n")

	{
		b.WriteString("\n")

		b.WriteString("func UserAgentBrowser")
		b.WriteString("() string {\n")
		b.WriteString("\treturn Choose([]string{\n")

		for _, value := range uaList {
			b.WriteString("\t\t")
			b.WriteString("`")
			b.WriteString(value.Useragent)
			b.WriteString("`")
			b.WriteString(",\n")
		}

		b.WriteString("\t})\n")
		b.WriteString("}\n")
	}

	for key, values := range m {
		b.WriteString("\n")

		b.WriteString("func UserAgent")
		b.WriteString(key)
		b.WriteString("() string {\n")
		b.WriteString("\treturn Choose([]string{\n")

		values = candy.UniqueUsing(values, func(ua *UA) any {
			return ua.Useragent
		})

		for _, value := range values {
			b.WriteString("\t\t")
			b.WriteString("`")
			b.WriteString(value.Useragent)
			b.WriteString("`")
			b.WriteString(",\n")
		}

		b.WriteString("\t})\n")
		b.WriteString("}\n")
	}

	err = os.WriteFile("user_agent_browser.go", b.Bytes(), 0600)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
}
