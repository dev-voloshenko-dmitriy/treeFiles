package cmd

import (
	"fmt"
	"os"

	"github.com/vdsEngineer/treeFiles/src/internal"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:  "treeFiles",
		Usage: "tree Files - minimalistic, fast and customizable cli applications for viewing files in the terminal in the form of a tree",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "enclosure",
				Usage:   "the nesting level for the tree",
				Aliases: []string{"e"},
				Value:   2,
			},
			&cli.IntFlag{
				Name:    "limitFiles",
				Usage:   "limit of files to show in the tree",
				Aliases: []string{"l"},
				Value:   15,
			},
			&cli.StringFlag{
				Name:    "rootFontColor",
				Usage:   "font color for the root of the tree",
				Aliases: []string{"rc"},
				Value:   "35",
			},
			&cli.StringFlag{
				Name:    "colorLine",
				Usage:   "color for the lines of the tree",
				Aliases: []string{"lc"},
				Value:   "63",
			},
			&cli.StringFlag{
				Name:    "itemFontColor",
				Usage:   "font color for the items of the tree",
				Aliases: []string{"ic"},
				Value:   "120",
			},
		},
		Action: func(c *cli.Context) error {

			enumeratorStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color(c.String("colorLine"))).
				MarginRight(1)

			rootStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(c.String("rootFontColor")))
			itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(c.String("itemFontColor")))

			enclosure := c.Int("enclosure")

			limitFiles := c.Int("limitFiles")

			treeFiles, err := internal.GetTree(".", enclosure, limitFiles, c.String("search"))
			if err != nil {
				fmt.Println(err)
			}

			treeFiles.
				Enumerator(tree.RoundedEnumerator).
				EnumeratorStyle(enumeratorStyle).
				RootStyle(rootStyle).
				ItemStyle(itemStyle)

			fmt.Println(treeFiles)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
