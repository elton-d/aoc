package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"runtime"
)

var (
	sessionCookie = ""
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	d := filepath.Join(filepath.Dir(b))

	bts, err := ioutil.ReadFile(filepath.Join(d, "cookie.txt"))
	if err != nil {
		panic(err)
	}
	sessionCookie = string(bts)
}

func GetInput(url string) ([]byte, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	path := filepath.Join("testdata", u.Path)
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
			res, err := (&http.Client{}).Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			if err := ioutil.WriteFile(path, b, os.ModePerm); err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	return ioutil.ReadFile(path)
}
