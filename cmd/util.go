package cmd

import (
	"context"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func handleCommand(f func(context.Context, *console.Client, *cobra.Command, []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var exitWithCode int
		defer func() {
			if exitWithCode > 0 {
				os.Exit(exitWithCode)
			}
		}()

		tele := telemetry.New(cmd.Context(), telemetryOpts()...)

		ctx, _ := tele.Tracer(cmd.Context())
		ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "cli:"+cmd.CommandPath(),
			trace.WithAttributes(
				attribute.String("command", cmd.CommandPath()),
			),
			trace.WithSpanKind(trace.SpanKindServer),
		)

		var exitError error
		defer func() {
			spanEnder(exitError, trace.WithStackTrace(false))
			tele.Shutdown(cmd.Context())
		}()

		if err := handleConsole(ctx, tele, cmd, args, f); err != nil {
			if ee, ok := err.(*console.ExitCodeError); ok {
				exitWithCode = ee.ExitCode
				return // exitError is not set
			}

			exitError, exitWithCode = err, 1
			cmd.PrintErrln(err)
		}
	}
}

func handleConsole(ctx context.Context, tele *telemetry.Client, cmd *cobra.Command, args []string, f func(context.Context, *console.Client, *cobra.Command, []string) error) error {
	configPath := viper.GetString("configPath")

	ctx, _ = signalcontext.WithInterrupt(ctx, logging.NewZHcLog(&log.Logger, ""))
	var c *console.Client

	delayMessage := ui.IsTerminal()

	switch cmd.Name() {
	// Don't init console client with these commands
	case "completion", "options":
		delayMessage = false
	case "init":
		// No console client created here
	default:
		var err error
		c, err = console.CreateClient(ctx, configPath)
		if err != nil {
			return err
		}
		if c.Client().HistoryCfg != nil {
			trace.SpanFromContext(ctx).SetAttributes(attribute.Bool("history_enabled", true))
		}

		defer c.Client().Close()
	}

	if tele.NewRandomId() {
		ui.ColorizedOutput(ui.ColorInfo, "Anonymous telemetry collection enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/cli/telemetry\n")
		if delayMessage {
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	if err := f(ctx, c, cmd, args); err != nil {
		return err
	}

	return nil
}