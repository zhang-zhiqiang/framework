package env

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/gohade/hade/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

type HadeEnv struct {
	folder string
	maps   map[string]string
}

func NewHadeEnv(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewHadeEnv param error")
	}

	folder := params[0].(string)

	hadeEnv := &HadeEnv{
		folder: folder,
		maps:   map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	file := path.Join(folder, ".env")

	fi, err := os.Open(file)
	if err == nil {
		defer fi.Close()

		br := bufio.NewReader(fi)
		for {
			line, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			s := bytes.SplitN(line, []byte{'='}, 2)
			if len(s) < 2 {
				continue
			}
			key := string(s[0])
			val := string(s[1])
			hadeEnv.maps[key] = val
		}
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		hadeEnv.maps[pair[0]] = pair[1]
	}

	return hadeEnv, nil
}

func (en *HadeEnv) AppEnv() string {
	return en.Get("APP_ENV")
}

func (en *HadeEnv) IsExist(key string) bool {
	_, ok := en.maps[key]
	return ok
}

func (en *HadeEnv) Get(key string) string {
	if val, ok := en.maps[key]; ok {
		return val
	}
	return ""
}

func (en *HadeEnv) All() map[string]string {
	return en.maps
}
