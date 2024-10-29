package api

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"fmt"
	"github.com/hjoshi123/temporal-loan-app/internal/config"
	"github.com/hjoshi123/temporal-loan-app/internal/database"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"github.com/hjoshi123/temporal-loan-app/internal/server"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Start the API server",
		RunE:  RunAPIServer,
	}

	disableMigration bool
	migrationPath    string
)

func Execute() error {
	apiCommand.Flags().StringVarP(&migrationPath, "migration-path", "m", "", "Path to the migration files")
	apiCommand.Flags().BoolVar(&disableMigration, "disable-migration", false, "Disable migration")
	return apiCommand.Execute()
}

func RunAPIServer(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	_ = database.Connect(ctx)

	if !disableMigration {
		workdir, err := atlasexec.NewWorkingDir(
			atlasexec.WithMigrations(
				os.DirFS("./migrations"),
			),
		)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to create working directory", "error", err)
			return err
		}
		// atlas exec works on a temporary directory, so we need to close it
		defer workdir.Close()

		// Initialize the client.
		client, err := atlasexec.NewClient(workdir.Path(), "atlas")
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to create atlas client", "error", err)
			return err
		}

		_, err = client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
			URL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Spec.DBUser, config.Spec.DBPassword, config.Spec.DBHost,
				config.Spec.DBPort, config.Spec.DBName),
		})
		if err != nil {
			logging.FromContext(ctx).Fatalw("failed to apply migrations", "error", err)
			return err
		}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	httpMux := server.Setup()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Spec.Port), c.Handler(httpMux)); err != nil {
		logging.FromContext(ctx).Fatalw("failed to start server", "error", err)
		return err
	}

	return nil
}
