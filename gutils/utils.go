package gutils

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
	"strconv"
)

// ReaderUnmarshal
// 读取reader中的内容,并使用json解析
func ReaderUnmarshal(reader io.Reader, info interface{}) error {
	bys, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}
	return json.Unmarshal(bys, info)
}

// ReaderMarshal
/*
使用json编码，并转换成reader
*/
func ReaderMarshal(src interface{}) (io.Reader, error) {
	buf, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}

func FileSizeFormat(size int64) (s_ string) {
	if s := float64(size) / 1024; s < 1 {
		s_ = strconv.FormatInt(size, 10) + " B"
	} else {
		if ss := s / 1024; ss < 1 {
			s_ = strconv.FormatFloat(s, 'f', 1, 32) + " K"
		} else {
			if sss := ss / 1024; sss < 1 {
				s_ = strconv.FormatFloat(ss, 'f', 1, 32) + " M"
			} else {
				if ssss := sss / 1024; ssss < 1 {
					s_ = strconv.FormatFloat(sss, 'f', 1, 32) + " G"
				} else {
					if sssss := ssss / 1024; sssss < 1 {
						s_ = strconv.FormatFloat(ssss, 'f', 1, 32) + " T"
					}
				}
			}
		}
	}

	return
}

// SlicePage 分页
func SlicePage(pageNum, pageSize, total int) (sliceStart int) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	pageCount := int(math.Ceil(float64(total) / float64(pageSize))) // 数据可以分为几页
	if pageNum > pageCount {                                        // 最后一页或超过最后一页,则返回最后一页
		pageNum = pageCount
	}
	sliceStart = (pageNum - 1) * pageSize

	return sliceStart
}
