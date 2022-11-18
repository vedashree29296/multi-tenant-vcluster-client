package cmd

import (
	"fmt"
	"mtc/pkg"

	"github.com/spf13/cobra"
)

// createCmd represents the `create` command.
var createTenantCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"create-tenant"},
	Short:   "Create a new tenant",
	Long:    "",
	RunE:    runCreateTenant,
}

func runCreateTenant(cmd *cobra.Command, args []string) error {
	tenantName := args[0]
	command := "vcluster"
	options := []string{"create", tenantName}
	fmt.Println(command)
	err := pkg.ExecCommand(command, options, false)
	if err != nil {
		pkg.ShowError(err.Error())
		return err
	}
	return nil
}
