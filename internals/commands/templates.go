package commands

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
    
    
)


// ProjectConfig holds the project configuration
type ProjectConfig struct {
    ProjectName string
    Language    string
    Port        int
    Version     string
}


// GenerateGoProject creates a Go-based project structure
func GenerateGoProject(projectName string, config ProjectConfig) error {
    files := map[string]string{
        "Dockerfile":        goDockerfileTemplate,
        "docker-compose.yml": dockerComposeTemplate,
        ".deploygo.yml":     deploygoConfigTemplate,
        "go.mod":           fmt.Sprintf("module %s\n\ngo 1.21\n", projectName),
        "main.go":          goMainTemplate,
        ".dockerignore":     dockerIgnoreTemplate,
    }
    
    return createFiles(projectName, files, config)
}

// GenerateNodeJSProject creates a Node.js project structure  
func GenerateNodeJSProject(projectName string, config ProjectConfig) error {
    files := map[string]string{
        "Dockerfile":        nodejsDockerfileTemplate,
        "docker-compose.yml": dockerComposeTemplate,
        ".deploygo.yml":     deploygoConfigTemplate,
        "package.json":      fmt.Sprintf(nodejsPackageTemplate, projectName),
        "app.js":           nodejsAppTemplate,
        ".dockerignore":     dockerIgnoreTemplate,
    }
    
    return createFiles(projectName, files, config)
}

// GeneratePythonProject creates a Python project structure
func GeneratePythonProject(projectName string, config ProjectConfig) error {
    files := map[string]string{
        "Dockerfile":        pythonDockerfileTemplate,
        "docker-compose.yml": dockerComposeTemplate,
        ".deploygo.yml":     deploygoConfigTemplate,
        "requirements.txt":  "flask==2.3.3\n",
        "app.py":           pythonAppTemplate,
        ".dockerignore":     dockerIgnoreTemplate,
    }
    
    return createFiles(projectName, files, config)
}

func createFiles(projectName string, files map[string]string, config ProjectConfig) error {
    for filename, content := range files {
        filePath := filepath.Join(projectName, filename)
        
        // Create directory if needed
        dir := filepath.Dir(filePath)
        if err := os.MkdirAll(dir, 0755); err != nil {
            return err
        }
        
        // Create and write file
        file, err := os.Create(filePath)
        if err != nil {
            return err
        }
        defer file.Close()
        
        // For templated content, execute the template
        tmpl, err := template.New(filename).Parse(content)
        if err != nil {
            // If template parsing fails, write raw content
            if _, err := file.WriteString(content); err != nil {
                return err
            }
        } else {
            // Execute template with config
            if err := tmpl.Execute(file, config); err != nil {
                return err
            }
        }
        
        fmt.Printf("📄 Created: %s\n", filePath)
    }
    return nil
}