package common

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/CrossChEp/kv-db/internal/entity"
)

func ParseBufferSize(bufferSize string) (int, error) {
	if bufferSize == "" {
		return 0, fmt.Errorf("%w: buffer size is empty", entity.ErrInvalidBufferSizeFormat)
	}

	var (
		bsStr  string
		format string
	)

	for i, el := range bufferSize {
		if !unicode.IsLetter(el) && !unicode.IsDigit(el) {
			return 0, entity.ErrInvalidBufferSizeFormat
		}

		if unicode.IsLetter(el) {
			bsStr = bufferSize[:i]
			format = bufferSize[i:]

			break
		}
	}

	if format == "" {
		format = "B"
		bsStr = bufferSize
	}

	if bsStr == "" {
		return 0, entity.ErrInvalidBufferSizeFormat
	}

	bs, err := strconv.Atoi(bsStr)
	if err != nil {
		return 0, err
	}

	format = strings.ToLower(format)

	if _, ok := entity.SuffixSizeFormatToSize[format]; !ok {
		return 0, entity.ErrInvalidBufferSizeSuffix
	}

	return bs * entity.SuffixSizeFormatToSize[format], nil
}
