package root

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "this is short dict",
	Long:  "this is A LONG DIct",

	Run: scan,
}

var _ = `
                          /\          /\
                         ( \\        // )
                          \ \\      // /
                           \_\\||||//_/
                            \/ _  _ \
                           \/|(O)(O)|
          ❤  bupd         \/ |      |
      ___________________\/  \      /
     //                //     |____|
    //                ||     /      \
   //|                \|     \ 0  0 /
  // \       )         V    / \____/
 //   \     /        (     /
""     \   /_________|  |_/
       /  /\   /     |  ||
      /  / /  /      \  ||
      | |  | |        | ||
      | |  | |        | ||
      |_|  |_|        |_||
       \_\  \_\        \_\\

  `

var logo2 = `
git-donkey

      \\__//
       /OO\\_______
       \__/\       )\/\
           ||----/ |
   ❤  bupd ||     ||

  `

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81")).Bold(true)

func init() {
	rootCmd.AddCommand(scanCmd)
}

func scan(cmd *cobra.Command, args []string) {
	fmt.Printf("Scanning through repos...\n")
	fmt.Printf("This may take a while Please wait...\n\n")

	fmt.Printf("%s\n", logoStyle.Render(logo2))
}
