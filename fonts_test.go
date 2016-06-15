package fonts
import (
	"testing"
	"github.com/golang/freetype/truetype"

)


func TestAppendPath(t *testing.T) {
	prevLen := len(FontPaths)
	AppendPath("abc")
	nowLen := len(FontPaths)
	if nowLen - prevLen != 1 {
		t.Error("AppendValidExt fail");
	}
}

func TestAppendValidExt(t *testing.T) {
	prevLen := len(ValidExt)
	AppendValidExt("abc")
	nowLen := len(ValidExt)
	if nowLen - prevLen != 1 {
		t.Error("AppendValidExt fail");
	}
}

func TestLoad(t *testing.T) {
	_, e := LoadFont("Microsoft Sans Serif");
	if e != nil {
		t.Error("not expected error");
	}
}

func TestTrueTypeLoad(t *testing.T) {
	data, _ := LoadFont("Times New Roman");
	// TTF or TTC
	_, e := truetype.Parse(data);
	if e != nil {
		t.Error(e);
		t.Error("not expected error");
	}
}


func TestLoadError(t *testing.T) {
	_, e := LoadFont("abc");
	if e == nil {
		t.Error("expected error");
	}
}