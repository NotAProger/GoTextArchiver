package cmd

import (
	"archiver/lib/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file useing virable-length code",
	Run:   unpack,
}

// TODO: take extention from file
const unpackedExtention = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		hanldeErr(ErrEmptyFile)
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

	unpacked := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(unpacked), 0644)
	if err != nil {
		hanldeErr(err)
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	//								//	path = /path/to/file/myFile.txt
	fileName := filepath.Base(path) // /path/to/file/myFile.txt -> myFile.txt
	//filepath.Ext(fileName)            // myFile.txt -> .txt
	//strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFile'

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtention
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
