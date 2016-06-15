package fonts
import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
)

func AppendPath(path string) {
	for _, p := range FontPaths {
		if p == path {
			return;
		}
	}
	FontPaths = append(FontPaths, path);
}

func AppendValidExt(ext string) {
	for _, ve := range ValidExt {
		if ve == ext {
			return;
		}
	}
	ValidExt = append(ValidExt, ext);
}

var ValidExt = []string{"ttf", "ttc"}

func LoadFont(fontName string) ([]byte, error) {
	for _, p := range FontPaths {
		for _, ext := range ValidExt {
			matches, err := filepath.Glob(os.ExpandEnv(p) + fontName + "." + ext)
			if err != nil {
				return nil, err
			}
			for _, match := range matches {
				return ioutil.ReadFile(match)
			}
		}
	}
	return nil, fmt.Errorf("Font '%s' not found.", fontName)
}

