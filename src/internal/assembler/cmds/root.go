package cmds

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hasanaburayyan/byte-machine/src/internal/assembler"
	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/spf13/cobra"
)

var outputFileName string
var inputFileName string
var run bool
var outSourceMap bool

var rootCmd = &cobra.Command{
	Use:   "bmasm",
	Short: "Compile bm assembly to byte code",
	Long:  "bmasm is a cli for compiling and running bm assembly code",
	Run: func(cmd *cobra.Command, args []string) {
		var r io.Reader

		if !run && outputFileName == "" {
			fmt.Println("error: --out is required unless --run is set")
			os.Exit(1)
		}

		if inputFileName == "" {
			stdinBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Printf("reading input from stdin: %v\n", err)
				os.Exit(1)
			}
			inputText := string(stdinBytes)
			r = strings.NewReader(inputText)
		} else {
			file, err := os.Open(inputFileName)
			if err != nil {
				fmt.Printf("error reading file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()

			r = bufio.NewReader(file)
		}

		res, sourcemap, err := assembler.Assemble(r)
		if err != nil {
			fmt.Printf("assembly error: %v\n", err)
			os.Exit(1)
		}

		if run {
			bm := bytemachine.NewByteMachine(res)
			bm.Run()
			return
		}

		file, err := os.Create(outputFileName)
		if err != nil {
			fmt.Printf("error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.Write(res)
		if err != nil {
			fmt.Printf("error writing bytes to file: %v\n", err)
			os.Exit(1)
		}

		if outSourceMap {
			f, err := os.Create(replaceExtension(outputFileName, ".bmsmap.json"))
			if err != nil {
				fmt.Printf("error creating output file: %v\n", err)
				os.Exit(1)
			}
			defer f.Close()

			d, err := json.Marshal(sourcemap)
			if err != nil {
				fmt.Printf("marshal sourcemap: %v\n", err)
				os.Exit(1)
			}
			_, err = f.Write(d)
			if err != nil {
				fmt.Printf("writing sourcemap: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func replaceExtension(filename, newExt string) string {
	return filepath.Base(filename[:len(filename)-len(filepath.Ext(filename))]) + newExt
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&outputFileName, "out", "o", "", "The name of the output file for the binary")
	rootCmd.PersistentFlags().StringVarP(&inputFileName, "in", "i", "", "The name of the input file containing assembly code")
	rootCmd.PersistentFlags().BoolVarP(&run, "run", "r", false, "Skips writing binary to a file and just runs assembly code")
	rootCmd.PersistentFlags().BoolVarP(&outSourceMap, "source-map", "s", false, "Print out source map")

	rootCmd.Execute()
}
