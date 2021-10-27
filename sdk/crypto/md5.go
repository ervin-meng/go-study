package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func StudyMd5(raw string) string {
	md5Obj := md5.New()

	_, _ = io.WriteString(md5Obj, raw)

	return hex.EncodeToString(md5Obj.Sum(nil))
}
