package fonts
import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
	"io"
)
// AppendPath append search path
func AppendPath(path string) {
	for _, p := range FontPaths {
		if p == path {
			return;
		}
	}
	FontPaths = append(FontPaths, path);
}
// AppendValidExt append match ext
func AppendValidExt(ext string) {
	for _, ve := range ValidExt {
		if ve == ext {
			return;
		}
	}
	ValidExt = append(ValidExt, ext);
}

var ValidExt = []string{"ttf"}
// LoadFont
func LoadFont(fontName string) ([]byte, error) {
	r, e := ReadFont(fontName)
	if e != nil {
		return nil, e
	}
	return ioutil.ReadAll(r);
}

func ReadFont(fontName string) (io.Reader, error) {
	for _, p := range FontPaths {
		for _, ext := range ValidExt {
			matches, err := filepath.Glob(os.ExpandEnv(p) + fontName + "." + ext)
			if err != nil {
				return nil, err
			}
			for _, match := range matches {
				return os.Open(match)
			}
		}
	}
	return nil, fmt.Errorf("Font '%s' not found.", fontName)
}

