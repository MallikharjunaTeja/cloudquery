package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
)

const policyRunHelpMsg = "Executes a policy on CloudQuery database"

var (
	policyRunCmd = &cobra.Command{
		Use:   "run",
		Short: policyRunHelpMsg,
		Long:  policyRunHelpMsg,
		Example: `
  # Download & Run the policies defined in your config
  cloudquery policy run

  # Run a specific policy by it's name
  cloudquery policy run --policy my_aws_policy

  # See https://hub.cloudquery.io for additional policies.`,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.RunPolicies(ctx, args, policyName, outputDir, stopOnFailure, skipVersioning, failOnViolation, noResults)
		}),
	}
	outputDir       string
	policyName      string
	stopOnFailure   bool
	skipVersioning  bool
	failOnViolation bool
	noResults       bool
)

func init() {
	flags := policyRunCmd.Flags()
	flags.StringVar(&policyName, "policy", "", "Select specific policy to run")
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&stopOnFailure, "stop-on-failure", false, "Stops the policy execution on the first failure")
	flags.BoolVar(&failOnViolation, "fail-on-violation", false, "Return non zero exit code if one of the policy is violated")
	flags.BoolVar(&skipVersioning, "skip-versioning", false, "Skip policy versioning and use latest files")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)
}