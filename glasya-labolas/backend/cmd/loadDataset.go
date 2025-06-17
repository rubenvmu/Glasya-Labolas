package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"glasya-labolas/parsers"
)

var loadDatasetCmd = &cobra.Command{
	Use:   "load-dataset",
	Short: "Carga un CSV de entidades",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Debes especificar un archivo CSV")
			return
		}
		filePath := args[0]
		if err := parsers.LoadCSV(filePath); err != nil {
			fmt.Println("❌ Error al cargar CSV:", err)
			return
		}
		fmt.Println("✅ Dataset cargado correctamente")
	},
}

func init() {
	rootCmd.AddCommand(loadDatasetCmd)
}