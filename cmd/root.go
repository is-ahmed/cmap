package cmd

import (
	"github.com/spf13/cobra"
)

var (
    RootCmd = &cobra.Command{
	 Use: "Command Map",
	 Short: "Map for common commands you use but not often enough to commit to memory",
    }
)



