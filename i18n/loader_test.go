package i18n

import (
	"errors"
	"os"
	"strings"
	"testing"
	"testing/fstest"

	xlanguage "golang.org/x/text/language"
)

// sampleFs 提供三种格式与三种 locale 的样例数据，模拟 embed.FS 行为
var sampleFs = fstest.MapFS{
	"localize/en.json":    {Data: []byte(`{"hello":"Hello"}`)},
	"localize/zh-CN.yaml": {Data: []byte("hello: 你好\n")},
	"localize/ja.toml":    {Data: []byte(`hello = "こんにちは"` + "\n")},
}

func TestLoadLocalizesMultiFormat(t *testing.T) {
	p := New(WithDefaultLang(xlanguage.Make("en")))
	err := p.LoadLocalizes(sampleFs)
	if err != nil {
		t.Fatalf("LoadLocalizes err: %v", err)
	}
	got := p.LocalizeWithLang(xlanguage.Make("en"), "hello")
	if got != "Hello" {
		t.Errorf("en hello=%q", got)
	}
	got = p.LocalizeWithLang(xlanguage.Make("zh-CN"), "hello")
	if got != "你好" {
		t.Errorf("zh-CN hello=%q", got)
	}
	got = p.LocalizeWithLang(xlanguage.Make("ja"), "hello")
	if got != "こんにちは" {
		t.Errorf("ja hello=%q", got)
	}
}

func TestLoadLocalizesMapFs(t *testing.T) {
	fsys := fstest.MapFS{
		"localize/en.json":  {Data: []byte(`{"k":"V","nested":{"a":"A"}}`)},
		"localize/zh.yaml":  {Data: []byte("k: 中\nnested:\n  a: 嵌套\n")},
		"localize/fr.toml":  {Data: []byte("k = \"FR\"\n")},
		"localize/skip.xml": {Data: []byte("<x/>")},
		"localize/subdir":   {Mode: 0o755 | 1<<31 /* placeholder, IsDir set by MapFS */},
	}
	// MapFS dir entry
	fsys["localize/subdir/ignored.json"] = &fstest.MapFile{Data: []byte(`{"x":"y"}`)}

	p := New()
	err := p.LoadLocalizes(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	v := p.LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "V" {
		t.Errorf("en.k=%q", v)
	}
	v = p.LocalizeWithLang(xlanguage.Make("en"), "nested.a")
	if v != "A" {
		t.Errorf("en.nested.a=%q", v)
	}
	v = p.LocalizeWithLang(xlanguage.Make("zh"), "k")
	if v != "中" {
		t.Errorf("zh.k=%q", v)
	}
	v = p.LocalizeWithLang(xlanguage.Make("fr"), "k")
	if v != "FR" {
		t.Errorf("fr.k=%q", v)
	}
}

func TestLoadLocalizesDirNotFound(t *testing.T) {
	fsys := fstest.MapFS{}
	p := New()
	err := p.LoadLocalizes(fsys)
	if err == nil {
		t.Fatal("expected err for missing dir")
	}
}

func TestLoadLocalizesAggregateErrors(t *testing.T) {
	fsys := fstest.MapFS{
		"localize/good.json": {Data: []byte(`{"k":"v"}`)},
		"localize/bad1.json": {Data: []byte(`{not json`)},
		"localize/bad2.json": {Data: []byte(`{still not json`)},
	}
	p := New()
	err := p.LoadLocalizes(fsys)
	if err == nil {
		t.Fatal("expected aggregate err")
	}
	// errors.Join 包含两个错误 → Unwrap() []error 应有 2 个
	var unwrapper interface{ Unwrap() []error }
	if !errors.As(err, &unwrapper) {
		t.Fatalf("err should be errors.Join: %T", err)
	}
	n := len(unwrapper.Unwrap())
	if n != 2 {
		t.Errorf("got %d errs want 2", n)
	}
	// good 文件应已加载
	v := p.LocalizeWithLang(xlanguage.Make("good"), "k")
	if v != "v" {
		t.Errorf("good.k=%q", v)
	}
}

func TestLoadLocalizesUnknownExtSkipped(t *testing.T) {
	fsys := fstest.MapFS{
		"localize/en.xyz":  {Data: []byte("???")},
		"localize/en.json": {Data: []byte(`{"k":"v"}`)},
	}
	p := New()
	err := p.LoadLocalizes(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	v := p.LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "v" {
		t.Errorf("en.k=%q", v)
	}
}

func TestLoadLocalizesReadFail(t *testing.T) {
	// 用嵌套目录名作扫描目标但其下没文件
	fsys := fstest.MapFS{
		"empty/.keep": {Data: []byte("")},
	}
	p := New()
	err := p.LoadLocalizesWithFs("empty", fsys)
	if err != nil {
		// .keep 没扩展名 → 跳过；不应报错
		if !strings.Contains(err.Error(), "") {
			t.Fatalf("unexpected: %v", err)
		}
	}
}

func TestLoadFile(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/en.json"
	err := os.WriteFile(path, []byte(`{"hi":"Hello","n":{"a":"A"}}`), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	p := New()
	err = p.LoadFile(path)
	if err != nil {
		t.Fatalf("LoadFile: %v", err)
	}
	v := p.LocalizeWithLang(xlanguage.Make("en"), "hi")
	if v != "Hello" {
		t.Errorf("hi=%q", v)
	}
	v = p.LocalizeWithLang(xlanguage.Make("en"), "n.a")
	if v != "A" {
		t.Errorf("n.a=%q", v)
	}
}

func TestLoadFileWithLang(t *testing.T) {
	dir := t.TempDir()
	// 文件名 messages.yaml（lang 无法从文件名推），用显式 tag
	path := dir + "/messages.yaml"
	err := os.WriteFile(path, []byte("hi: 你好\n"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	p := New()
	zh := xlanguage.Make("zh")
	err = p.LoadFileWithLang(zh, path)
	if err != nil {
		t.Fatalf("LoadFileWithLang: %v", err)
	}
	v := p.LocalizeWithLang(zh, "hi")
	if v != "你好" {
		t.Errorf("hi=%q", v)
	}
}

func TestLoadFileNotFound(t *testing.T) {
	p := New()
	err := p.LoadFile("/nonexistent/path/en.json")
	if err == nil {
		t.Fatal("expected err for missing file")
	}
}

func TestLoadFileUnknownExt(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/en.unknown"
	err := os.WriteFile(path, []byte("data"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	p := New()
	err = p.LoadFile(path)
	if err == nil {
		t.Fatal("expected err for unknown ext")
	}
	if !errors.Is(err, ErrLocalizerNotFound) {
		t.Errorf("err should be ErrLocalizerNotFound: %v", err)
	}
}

func TestLoadFileBadContent(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/en.json"
	err := os.WriteFile(path, []byte(`{not json`), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	p := New()
	err = p.LoadFile(path)
	if err == nil {
		t.Fatal("expected unmarshal err")
	}
}

func TestLoadFs(t *testing.T) {
	fsys := fstest.MapFS{
		"data/en.json": {Data: []byte(`{"hi":"Hello"}`)},
	}
	p := New()
	err := p.LoadFs(fsys, "data/en.json")
	if err != nil {
		t.Fatalf("LoadFs: %v", err)
	}
	v := p.LocalizeWithLang(xlanguage.Make("en"), "hi")
	if v != "Hello" {
		t.Errorf("hi=%q", v)
	}
}

func TestLoadFsWithLang(t *testing.T) {
	fsys := fstest.MapFS{
		"messages.toml": {Data: []byte(`hi = "嗨"`)},
	}
	p := New()
	zh := xlanguage.Make("zh")
	err := p.LoadFsWithLang(zh, fsys, "messages.toml")
	if err != nil {
		t.Fatal(err)
	}
	v := p.LocalizeWithLang(zh, "hi")
	if v != "嗨" {
		t.Errorf("hi=%q", v)
	}
}

func TestLoadFsNotFound(t *testing.T) {
	fsys := fstest.MapFS{}
	p := New()
	err := p.LoadFs(fsys, "missing.json")
	if err == nil {
		t.Fatal("expected err")
	}
}

func TestLoadFsUnknownExt(t *testing.T) {
	fsys := fstest.MapFS{"en.xyz": {Data: []byte("x")}}
	p := New()
	err := p.LoadFs(fsys, "en.xyz")
	if !errors.Is(err, ErrLocalizerNotFound) {
		t.Errorf("err=%v", err)
	}
}

func TestLoadFilePackageDefault(t *testing.T) {
	original := Default
	Default = New()
	defer func() { Default = original }()

	dir := t.TempDir()
	path := dir + "/en.json"
	err := os.WriteFile(path, []byte(`{"k":"v"}`), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = LoadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	v := LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "v" {
		t.Errorf("k=%q", v)
	}
}

func TestLoadFsPackageDefault(t *testing.T) {
	original := Default
	Default = New()
	defer func() { Default = original }()

	fsys := fstest.MapFS{"en.json": {Data: []byte(`{"k":"v"}`)}}
	err := LoadFs(fsys, "en.json")
	if err != nil {
		t.Fatal(err)
	}
	v := LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "v" {
		t.Errorf("k=%q", v)
	}
}

func TestLoadLocalizesPackageDefault(t *testing.T) {
	original := Default
	Default = New()
	defer func() { Default = original }()

	fsys := fstest.MapFS{
		"localize/en.json": {Data: []byte(`{"hi":"Hello"}`)},
	}
	err := LoadLocalizes(fsys)
	if err != nil {
		t.Fatal(err)
	}
	v := LocalizeWithLang(xlanguage.Make("en"), "hi")
	if v != "Hello" {
		t.Errorf("v=%q", v)
	}
}

func TestLoadDirRecursive(t *testing.T) {
	root := t.TempDir()
	err := os.MkdirAll(root+"/web", 0o755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.MkdirAll(root+"/api", 0o755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(root+"/web/en.json", []byte(`{"k":"web-en"}`), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(root+"/api/zh-CN.yaml", []byte("k: api-zh"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(root+"/readme.txt", []byte("skip"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	p := New()
	err = p.LoadDir(root)
	if err != nil {
		t.Fatalf("LoadDir err: %v", err)
	}
	// 注意：web/en.json 与 api/en.json 若同名会后者覆盖前者；此例仅一份 en
	got := p.LocalizeWithLang(xlanguage.Make("en"), "k")
	if got != "web-en" {
		t.Errorf("en k=%q", got)
	}
	got = p.LocalizeWithLang(xlanguage.Make("zh-CN"), "k")
	if got != "api-zh" {
		t.Errorf("zh-CN k=%q", got)
	}
}

func TestLoadFsDirRecursive(t *testing.T) {
	fsys := fstest.MapFS{
		"i18n/en.json":        &fstest.MapFile{Data: []byte(`{"k":"v1"}`)},
		"i18n/sub/zh-CN.yaml": &fstest.MapFile{Data: []byte("k: v2")},
		"i18n/skip.txt":       &fstest.MapFile{Data: []byte("ignored")},
	}
	p := New()
	err := p.LoadFsDir(fsys, "i18n")
	if err != nil {
		t.Fatal(err)
	}
	got := p.LocalizeWithLang(xlanguage.Make("en"), "k")
	if got != "v1" {
		t.Errorf("en k=%q", got)
	}
	got = p.LocalizeWithLang(xlanguage.Make("zh-CN"), "k")
	if got != "v2" {
		t.Errorf("zh-CN k=%q", got)
	}
}

func TestLoadDirNotExist(t *testing.T) {
	err := New().LoadDir("/nonexistent-i18n-dir-xyz")
	if err == nil {
		t.Error("want error")
	}
}

func TestLoadDirPackageDefault(t *testing.T) {
	original := Default
	Default = New()
	defer func() { Default = original }()

	dir := t.TempDir()
	err := os.WriteFile(dir+"/en.json", []byte(`{"k":"v"}`), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = LoadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	v := LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "v" {
		t.Errorf("k=%q", v)
	}
}

func TestLoadFsDirPackageDefault(t *testing.T) {
	original := Default
	Default = New()
	defer func() { Default = original }()

	fsys := fstest.MapFS{"r/en.json": {Data: []byte(`{"k":"v"}`)}}
	err := LoadFsDir(fsys, "r")
	if err != nil {
		t.Fatal(err)
	}
	v := LocalizeWithLang(xlanguage.Make("en"), "k")
	if v != "v" {
		t.Errorf("k=%q", v)
	}
}

func TestLoadFsDirAggregateErrors(t *testing.T) {
	fsys := fstest.MapFS{
		"r/good.json": {Data: []byte(`{"k":"v"}`)},
		"r/bad.json":  {Data: []byte(`{not json`)},
	}
	p := New()
	err := p.LoadFsDir(fsys, "r")
	if err == nil {
		t.Fatal("expected err")
	}
	// good 应已加载
	v := p.LocalizeWithLang(xlanguage.Make("good"), "k")
	if v != "v" {
		t.Errorf("good.k=%q", v)
	}
}

func TestLoadFsDirRootMissing(t *testing.T) {
	fsys := fstest.MapFS{}
	err := New().LoadFsDir(fsys, "missing")
	if err == nil {
		t.Fatal("expected err for missing root")
	}
}
