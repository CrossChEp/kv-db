package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/CrossChEp/kv-db/internal/network/client"
	"github.com/CrossChEp/kv-db/internal/utils/pointer"
)

const (
	defaultAddress    = "localhost:5000"
	defaultBufferSize = 4096
)

var (
	flagAddress    = flag.String("address", defaultAddress, "host address of database")
	flagBufferSize = flag.Int("buffer", defaultBufferSize, "size of message")
)

func main() {
	flag.Parse()

	var (
		addr     = defaultAddress
		buffSize = defaultBufferSize
	)

	if pointer.Deref(flagAddress) != defaultAddress {
		addr = *flagAddress
	}
	if pointer.Deref(flagBufferSize) != defaultBufferSize {
		buffSize = *flagBufferSize
	}

	cl, err := client.New(
		addr,
		client.WithBufferSize(buffSize),
	)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	writerErr := bufio.NewWriter(os.Stderr)

	for {
		query, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(fmt.Errorf("could not read query, err: %w", err))

			continue
		}

		res, err := cl.Send(query)
		if err != nil {
			fmt.Println(fmt.Errorf("could not send query to database, err: %w", err))
		}

		if strings.HasPrefix(res, "[err]") {
			writerErr.WriteString(res + "\n")
			writerErr.Flush()

			continue
		}

		writer.WriteString(res + "\n")
		writer.Flush()
	}
}
