package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "vibecheck",
	Short: "Generate PR text for vibe-coded projects",
	Long:  "A CLI tool to generate GitHub PR body text based on vibe coding metrics",
	Run:   runVibeCheck,
}

func init() {
	rootCmd.Flags().Int("vibes", 0, "Vibe percentage (0-100) [REQUIRED]")
	rootCmd.Flags().Int("ai", 0, "AI percentage (0-100)")
	rootCmd.Flags().String("agent", "ai", "AI agent used (amp, claude, cursor)")
	rootCmd.Flags().String("artifact", "unknown", "Artifact identifier")

	rootCmd.MarkFlagRequired("vibes")

	viper.BindPFlag("vibes", rootCmd.Flags().Lookup("vibes"))
	viper.BindPFlag("ai", rootCmd.Flags().Lookup("ai"))
	viper.BindPFlag("agent", rootCmd.Flags().Lookup("agent"))
	viper.BindPFlag("artifact", rootCmd.Flags().Lookup("artifact"))
}

func runVibeCheck(cmd *cobra.Command, args []string) {
	vibes := viper.GetInt("vibes")
	ai := viper.GetInt("ai")
	agent := viper.GetString("agent")
	artifact := viper.GetString("artifact")

	human := 100 - ai

	output := fmt.Sprintf("Sick! You wrote %d%% of this code, %s wrote %d%%, you're saying this is %d%% vibe coded, and the development artifacts are at url example.com/%s",
		human, agent, ai, vibes, artifact)

	fmt.Println(output)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
