package generateids

import (
	"strconv"

	"github.com/dewadg/go-playground-api/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "generate:ids",
		Run: func(cmd *cobra.Command, args []string) {
			var numOfIDs int
			var length int
			if len(args) >= 2 {
				val, err := strconv.Atoi(args[0])
				if err != nil {
					logrus.WithError(err).Fatal("failed to parse argument")
				}
				numOfIDs = val

				val, err = strconv.Atoi(args[1])
				if err != nil {
					logrus.WithError(err).Fatal("failed to parse argument")
				}
				length = val
			}

			if err := store.SeedIDs(cmd.Context(), numOfIDs, length); err != nil {
				logrus.WithError(err).Fatal("failed to execute generate:ids")
			}
		},
	}
}
