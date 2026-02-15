package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Long: `Generate shell completion scripts for recotem CLI.

To load completions:

Bash:
  $ source <(recotem completion bash)
  # To load completions for each session, execute once:
  # Linux:
  $ recotem completion bash > /etc/bash_completion.d/recotem
  # macOS:
  $ recotem completion bash > $(brew --prefix)/etc/bash_completion.d/recotem

Zsh:
  $ source <(recotem completion zsh)
  # To load completions for each session, execute once:
  $ recotem completion zsh > "${fpath[1]}/_recotem"

Fish:
  $ recotem completion fish | source
  # To load completions for each session, execute once:
  $ recotem completion fish > ~/.config/fish/completions/recotem.fish

PowerShell:
  PS> recotem completion powershell | Out-String | Invoke-Expression
`,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				return cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				return cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
			return nil
		},
	}
	return cmd
}
