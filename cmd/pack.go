package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file useing variable-length code",
	Run:   pack,
}

const packedExtention = "vlc"

var ErrEmptyFile = errors.New("path to file in not specified")

func pack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		hanldeErr(ErrEmptyFile)
	}

	var encoder compression.Encoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown comression method")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		hanldeErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		hanldeErr(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		hanldeErr(err)
	}
}

func packedFileName(path string) string {
	//								//	path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // /path/to/file/myFile.txt -> myFile.txt
	//filepath.Ext(fileName)            // myFile.txt -> .txt
	//strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFile'

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtention
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method choose: vlc")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic("method \"method\" have to be" + err.Error())
	}
}
