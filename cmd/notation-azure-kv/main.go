package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Azure/notation-azure-kv/internal/version"
	"github.com/notaryproject/notation-go/plugin/proto"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "notation-azure-kv",
		Usage:   "Notation - Notary V2 Azure KV plugin",
		Version: version.GetVersion(),
		Commands: []*cli.Command{
			metadataCommand,
			signCommand,
			describeKeyCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		var reer proto.RequestError
		if !errors.As(err, &reer) {
			err = proto.RequestError{
				Code: proto.ErrorCodeGeneric,
				Err:  err,
			}
		}
		data, _ := json.Marshal(err)
		os.Stderr.Write(data)
		os.Exit(1)
	}
}
