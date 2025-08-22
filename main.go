package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed cursor_logo.b64
var cursorLogo string

//go:embed amp_logo.b64
var ampLogo string

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
	rootCmd.Flags().String("output", "text", "Output format (text, json, github-markdown)")

	rootCmd.MarkFlagRequired("vibes")

	viper.BindPFlag("vibes", rootCmd.Flags().Lookup("vibes"))
	viper.BindPFlag("ai", rootCmd.Flags().Lookup("ai"))
	viper.BindPFlag("agent", rootCmd.Flags().Lookup("agent"))
	viper.BindPFlag("artifact", rootCmd.Flags().Lookup("artifact"))
	viper.BindPFlag("output", rootCmd.Flags().Lookup("output"))
}

type VibeData struct {
	Vibes    int    `json:"vibes"`
	AI       int    `json:"ai"`
	Human    int    `json:"human"`
	Agent    string `json:"agent"`
	Artifact string `json:"artifact"`
}

func runVibeCheck(cmd *cobra.Command, args []string) {
	vibes := viper.GetInt("vibes")
	ai := viper.GetInt("ai")
	agent := viper.GetString("agent")
	artifact := viper.GetString("artifact")
	outputFormat := viper.GetString("output")

	human := 100 - ai

	data := VibeData{
		Vibes:    vibes,
		AI:       ai,
		Human:    human,
		Agent:    agent,
		Artifact: artifact,
	}

	switch outputFormat {
	case "json":
		outputJSON(data)
	case "github-markdown":
		outputGitHubMarkdown(data)
	default:
		outputText(data)
	}
}

func outputText(data VibeData) {
	artifactURL := generateArtifactURL(data.Artifact, data.Agent)
	output := fmt.Sprintf("Sick! You wrote %d%% of this code, %s wrote %d%%, you're saying this is %d%% vibe coded, and the development artifacts are at url %s",
		data.Human, data.Agent, data.AI, data.Vibes, artifactURL)
	fmt.Println(output)
}

func outputJSON(data VibeData) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func outputGitHubMarkdown(data VibeData) {
	// Generate badges
	vibeBadge := generateVibeBadge(data.Vibes)
	aiBadge := generateAIBadge(data.AI, data.Agent)
	if data.Artifact != "unknown" {
		artifactBadge := generateArtifactBadge(data.Artifact, data.Agent)
		fmt.Printf("%s %s %s\n", vibeBadge, aiBadge, artifactBadge)
	} else {
		fmt.Printf("%s %s\n", vibeBadge, aiBadge)
	}
}

func generateVibeBadge(vibes int) string {
	color := "white"
	if vibes >= 75 {
		color = "green"
	} else if vibes >= 50 {
		color = "yellow"
	} else if vibes >= 25 {
		color = "orange"
	}
	return fmt.Sprintf("![vibe](https://img.shields.io/badge/vibe-%d%%25-%s)", vibes, color)
}

func generateAIBadge(ai int, agent string) string {
	color := "white"
	if ai >= 75 {
		color = "red"
	} else if ai >= 50 {
		color = "orange"
	} else if ai >= 25 {
		color = "yellow"
	}

	var logo string
	switch strings.ToLower(agent) {
	case "cursor":
		logo = url.QueryEscape(cursorLogo)
	case "amp":
		logo = url.QueryEscape(ampLogo)
	case "claude":
		logo = "anthropic"
	default:
		logo = "" // no logo for unknown agents
	}

	if logo != "" {
		if logo == "anthropic" {
			return fmt.Sprintf("![AI](https://img.shields.io/badge/AI-%d%%25-%s?logo=%s)", ai, color, logo)
		} else {
			return fmt.Sprintf("![AI](https://img.shields.io/badge/AI-%d%%25-%s?logo=data:image/svg%%2bxml;base64,%s)", ai, color, logo)
		}
	}
	return fmt.Sprintf("![AI](https://img.shields.io/badge/AI-%d%%25-%s)", ai, color)
}

func generateArtifactURL(artifact, agent string) string {
	if strings.ToLower(agent) == "amp" && strings.HasPrefix(artifact, "T-") {
		return fmt.Sprintf("https://ampcode.com/threads/%s", artifact)
	}
	return fmt.Sprintf("example.com/%s", artifact)
}

func escapeForShields(text string) string {
	// For shields.io, single dashes need to be escaped as double dashes
	return strings.ReplaceAll(text, "-", "--")
}

func generateArtifactBadge(artifact, agent string) string {
	if strings.ToLower(agent) == "amp" && strings.HasPrefix(artifact, "T-") {
		logo := url.QueryEscape(ampLogo)
		artifactURL := generateArtifactURL(artifact, agent)
		escapedArtifact := escapeForShields(artifact)
		return fmt.Sprintf("[![amp thread](https://img.shields.io/badge/amp_thread-%s-green?logo=data:image/svg%%2bxml;base64,%s)](%s)", escapedArtifact, logo, artifactURL)
	}
	
	var logo string
	escapedArtifact := escapeForShields(artifact)
	switch strings.ToLower(agent) {
	case "amp":
		logo = url.QueryEscape(ampLogo)
		return fmt.Sprintf("![artifact](https://img.shields.io/badge/artifact-%s-blue?logo=data:image/svg%%2bxml;base64,%s)", escapedArtifact, logo)
	default:
		return fmt.Sprintf("![artifact](https://img.shields.io/badge/artifact-%s-blue)", escapedArtifact)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
