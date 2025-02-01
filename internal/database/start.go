package database

import (
	"bufio"
	"context"
	"os"
	"strings"
)

func (db *Database) Start() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	writerErr := bufio.NewWriter(os.Stderr)

	ctx := context.Background()

	for {
		query, err := reader.ReadString('\n')
		if err != nil {
			db.log.Error(ctx, "Could not ")
			continue
		}

		res := db.Handle(ctx, query)

		if strings.HasPrefix(res, "[err]") {
			writerErr.WriteString(res + "\n")
			writerErr.Flush()

			continue
		}

		writer.WriteString(res + "\n")
		writer.Flush()
	}
}
