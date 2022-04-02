package generateids

import (
	"github.com/dewadg/go-playground-api/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "generate:ids",
		Run: func(cmd *cobra.Command, args []string) {
			if err := store.SeedIDs(cmd.Context()); err != nil {
				logrus.WithError(err).Fatal("failed to execute generate:ids")
			}
		},
	}
}
