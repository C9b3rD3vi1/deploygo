package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
    
    "github.com/spf13/cobra"
    "github.com/C9b3rD3vi1/deploygo/internals/commands"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
    Use:   "init [project-name]",
    Short: "Initialize a new project with deployment configuration",
    Long: `Initialize creates a new project with optimized Dockerfile, 
docker-compose.yml, and deployment configuration.`,
    Args: cobra.ExactArgs(1),
    Run:   runInit,
}


func init() {
    // Add flags to init command
    initCmd.Flags().StringP("template", "t", "go", "Project template (go, nodejs, python)")
    initCmd.Flags().StringP("language", "l", "go", "Primary language")
    initCmd.Flags().IntP("port", "p", 8080, "Application port")
    initCmd.Flags().Bool("overwrite", false, "Overwrite existing files")
}

func runInit(cmd *cobra.Command, args []string) {
    projectName := args[0]
    
    // Get flag values
    templateType, _ := cmd.Flags().GetString("template")
    language, _ := cmd.Flags().GetString("language")
    port, _ := cmd.Flags().GetInt("port")
    overwrite, _ := cmd.Flags().GetBool("overwrite")
    
    // Create project config
    config := commands.ProjectConfig{
        ProjectName: projectName,
        Language:    language,
        Port:        port,
        Version:     "1.0.0",
    }
    
    // Create project directory
    if err := createProjectDirectory(projectName, overwrite); err != nil {
        fmt.Printf("Error creating project directory: %v\n", err)
        return
    }
    
    // Generate files based on template
    if err := generateProjectFiles(projectName, templateType, config); err != nil {
        fmt.Printf("Error generating project files: %v\n", err)
        return
    }
    
    fmt.Printf("✅ Successfully created project: %s\n", projectName)
    fmt.Printf("📁 Directory: %s\n", projectName)
    fmt.Printf("🐳 Language: %s\n", language)
    fmt.Printf("🚀 Next steps:\n")
    fmt.Printf("   cd %s\n", projectName)
    fmt.Printf("   deploygo build\n")
    fmt.Printf("   deploygo deploy staging\n")
}

//
func createProjectDirectory(projectName string, overwrite bool) error {
    // Check if directory exists
    if _, err := os.Stat(projectName); err == nil {
        if !overwrite {
            return fmt.Errorf("directory '%s' already exists. Use --overwrite to replace", projectName)
        }
        // Remove existing directory if overwrite is enabled
        if err := os.RemoveAll(projectName); err != nil {
            return fmt.Errorf("failed to remove existing directory: %v", err)
        }
    }
    
    // Create project directory
    if err := os.MkdirAll(projectName, 0755); err != nil {
        return fmt.Errorf("failed to create project directory: %v", err)
    }
    
    return nil
}

// 
func generateProjectFiles(projectName, templateType string, config commands.ProjectConfig) error {
    templates := map[string]func(string, commands.ProjectConfig) error{
        "go":     commands.GenerateGoProject,
        "nodejs": commands.GenerateNodeJSProject,
        "python": commands.GeneratePythonProject,
    }
    
    generator, exists := templates[templateType]
    if !exists {
        return fmt.Errorf("unsupported template type: %s", templateType)
    }
    
    return generator(projectName, config)
}