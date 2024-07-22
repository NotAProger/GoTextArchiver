package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"archiver/lib/compression/vlc/table/shenonFano"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file useing variable-length code",
	Run:   unpack,
}

func unpack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		hanldeErr(ErrEmptyFile)
	}

	var decoder compression.Decoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "sf":
		decoder = vlc.New(shenonFano.Generator{})
	default:
		cmd.PrintErr("unknown decomression method")
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

	unpacked := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(unpacked), 0644)
	if err != nil {
		hanldeErr(err)
	}
}

func unpackedFileName(path string) string {
	//								//	path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // /path/to/file/myFile.txt -> myFile.txt
	//filepath.Ext(fileName)            // myFile.txt -> .txt
	//strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFile'
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method choose: vlc")
	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic("method \"method\" have to be" + err.Error())
	}
}
