package util

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	sessionCookie = ""
	utilsDir      = ""
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	utilsDir = filepath.Join(filepath.Dir(b))

	// CI is set by GitHub Actions https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables#default-environment-variables
	if os.Getenv("CI") != "true" {
		bts, err := os.ReadFile(filepath.Join(utilsDir, "cookie.txt"))
		if err != nil {
			log.Default().SetOutput(os.Stderr)
			log.Printf("could not read session cookie %v", err)
		}
		sessionCookie = strings.TrimSpace(string(bts))
	}

}

func GetInput(url string) ([]byte, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	path := filepath.Join(filepath.Dir(utilsDir), "testdata", u.Path)
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return nil, err
			}
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Add("cookie", sessionCookie)
			if err != nil {
				return nil, err
			}
			fmt.Printf("making request: %v\n", req)
			res, err := (&http.Client{}).Do(req)
			if err != nil {
				return nil, err
			}
			if res.StatusCode < 200 || res.StatusCode >= 300 {
				return nil, fmt.Errorf("non-OK status code returned: %s", res.Status)
			}
			defer res.Body.Close()
			b, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			if err := os.WriteFile(path, b, os.ModePerm); err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	return os.ReadFile(path)
}

func GetInputStr(url string) string {
	b, err := GetInput(url)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
