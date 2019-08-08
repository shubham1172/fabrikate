package cmd

import (
	"errors"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/microsoft/fabrikate/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func unset(keys []string, environment, subcomponent string, removeSubcomponent bool) (err error) {
	// Load config
	componentConfig := core.NewComponentConfig(".")

	// Split component path delimited on "."
	subcomponentPath := []string{}
	if len(subcomponent) > 0 {
		subcomponentPath = strings.Split(subcomponent, ".")
	}

	// Split key paths delimited on "."
	keyPaths := [][]string{}
	for _, keyString := range keys {
		keyParts := strings.Split(keyString, ".")
		keyPaths = append(keyPaths, keyParts)
	}

	// Load target env config
	if err := componentConfig.Load(environment); err != nil {
		return err
	}

	// Remove the entire target component config if specified
	if removeSubcomponent {
		log.Info(emoji.Sprintf(""))
		if err = componentConfig.RemoveComponentConfig(subcomponentPath); err != nil {
			return err
		}
	} else {
		// Remove all keys from the config
		for _, keyPath := range keyPaths {
			if err = componentConfig.UnsetConfig(subcomponentPath, keyPath); err != nil {
				return err
			}
		}
	}

	// Write out the config
	return componentConfig.Write(environment)
}

type unsetCmdOpts struct {
	subcomponent       string
	environment        string
	removeSubcomponent bool
}

func newUnsetCmd() *cobra.Command {
	opts := &unsetCmdOpts{}

	cmd := &cobra.Command{
		Use:   "unset <config> [--subcomponent subcomponent] <path1> <path2> ...",
		Short: "Unsets a config value for a component for a particular config environment in the Fabrikate definition; deleting the key form the config.",
		Long: `Unsets a config value for a component for a particular config environment in the Fabrikate definition; deleting the key from the config.
eg.
$ fab unset --environment prod data.replicas username

Unsets the key of 'data.replicas' and 'username' in the 'prod' config for the current component.

$ fab unset --subcomponent "myapp" endpoint

Unsets the key of 'endpoint' in the 'common' config (the default) for subcomponent 'myapp'.

$ fab unset --subcomponent "myapp.mysubcomponent" data.replicas 

Unsets the subkey "replicas" in the key 'data' in the 'common' config (the default) for the subcomponent 'mysubcomponent' of the subcomponent 'myapp'.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 && inputFile == "" {
				return errors.New("'unset' takes a config path as the first parameter and one or more keys to remove thereafter")
			}

			removeComponent := cmd.Flag("remove-component").Value.String() == "true"
			return unset(args, opts.environment, opts.subcomponent, removeComponent)
		},
	}

	cmd.PersistentFlags().StringVar(&opts.environment, "environment", "common", "Environment this configuration should be removed from")
	cmd.PersistentFlags().StringVar(&opts.subcomponent, "subcomponent", "", "Subcomponent this configuration should be removed from")
	cmd.PersistentFlags().Bool("remove-component", false, "Remove the component config specified in --subcomponent entirely")

	return cmd
}

func init() {
	rootCmd.AddCommand(newUnsetCmd())
}
