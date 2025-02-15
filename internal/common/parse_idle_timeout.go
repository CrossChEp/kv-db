package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/CrossChEp/kv-db/internal/entity"
)

func ParseIdleTimeout(idleTimout string) (time.Duration, error) {
	if idleTimout == "" {
		return 0, fmt.Errorf("%w: idle timeout is empty", entity.ErrInvalidTimeoutFormat)
	}

	if len(idleTimout) <= 1 {
		return 0, entity.ErrInvalidTimeoutFormat
	}

	var (
		itStr  string
		format string
	)

	for i, el := range idleTimout {
		if !unicode.IsLetter(el) && !unicode.IsDigit(el) {
			return 0, entity.ErrInvalidTimeoutFormat
		}

		if unicode.IsLetter(el) {
			itStr = idleTimout[:i]
			format = idleTimout[i:]

			break
		}
	}

	if itStr == "" || format == "" {
		return 0, entity.ErrInvalidTimeoutFormat
	}

	format = strings.ToLower(format)

	if _, ok := entity.SuffixTimeFormatToTime[format]; !ok {
		return 0, entity.ErrInvalidTimeSuffixFormat
	}

	it, err := strconv.Atoi(itStr)
	if err != nil {
		return 0, err
	}

	return time.Duration(it) * entity.SuffixTimeFormatToTime[format], nil
}
