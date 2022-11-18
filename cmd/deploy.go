package cmd

import (
	"errors"
	"mtc/pkg"

	"github.com/spf13/cobra"
)

// createCmd represents the `create` command.
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"create-tenant"},
	Short:   "deploy an application",
	Long:    "",
	RunE:    runDeploy,
}

// mtc --tenant tenant1 --namespace my-namespace --yaml-file deployment.yaml

func runDeploy(cmd *cobra.Command, args []string) error {

	tenantName, _ := cmd.Flags().GetString("tenant")
	if tenantName == "" {
		err := errors.New("Please specify tenant")
		pkg.ShowError(err.Error())
		return err
	}
	//connect to vcluster tenant
	err := pkg.ExecCommand("vcluster", []string{"connect", tenantName}, false)
	pkg.ShowMessage("info", "Connecting to cluster "+tenantName, false, false)
	if err != nil {
		pkg.ShowError(err.Error())
		return err
	}
	// create a namespace
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace != "default" {
		pkg.ShowMessage("info", "Creating namespace "+namespace, false, false)
		err := pkg.ExecCommand("kubectl", []string{"create", "ns", namespace}, false)
		if err != nil {
			pkg.ShowError(err.Error())
			return err
		}
	}
	// deploy the application pod
	deploymentFile, _ := cmd.Flags().GetString("yaml-file")
	pkg.ShowMessage("info", "Deploying yaml file "+deploymentFile, false, false)
	if deploymentFile != "" {
		err := pkg.ExecCommand("kubectl", []string{"apply", "-f", deploymentFile, "-n", namespace}, false)
		if err != nil {
			pkg.ShowError(err.Error())
			return err
		}
	}
	// vcluster disconnect
	err = pkg.ExecCommand("vcluster", []string{"disconnect"}, false)
	pkg.ShowMessage("info", "Disconnecting from vcluster tenant "+tenantName, false, false)
	if err != nil {
		pkg.ShowError(err.Error())
		return err
	}
	pkg.ShowMessage("info", "Deployed", false, false)
	return nil
}

func init() {
	deployCmd.Flags().String("tenant", "", "Name of the tenant")
	deployCmd.Flags().String("yaml-file", "", "Yaml manifest file for kubernetes")
	deployCmd.Flags().String("namespace", "default", "Create a new namespace for deployment")
}
