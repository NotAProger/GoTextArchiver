package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file useing virable-length code",
	Run:   pack,
}

const packedExtention = "vlc"

var ErrEmptyFile = errors.New("path to file in not specified")

func pack(_ *cobra.Command, args []string) {
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

	packed := vlc.Encode(string(data))

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
	packCmd.AddCommand(vlcPackCmd)
}
