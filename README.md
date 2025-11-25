# AutoCommenter

AutoCommenter is a command-line interface (CLI) tool designed to automate documentation for Go projects. It leverages AI to scan your codebase, understand its structure, and generate code comments and `README.md` files.

## Overview

The tool operates by first analyzing your project files to build a comprehensive context. This context is then used by the configured AI provider (e.g., Gemini) to generate relevant and accurate documentation, helping to maintain code clarity and project understanding.

Key functionalities include:
*   **AI Provider Management:** Configure and switch between supported AI providers.
*   **Project Context Generation:** Scan the project directory to create a detailed context for the AI.
*   **Automated Commenting:** Generate comments for Go source files based on their content and the project context.
*   **README Generation:** Create a `README.md` file summarizing the project.

## Installation

To install the `ac` CLI tool, use the following `go install` command:

```bash
go install github.com/praneeth-ayla/AutoCommenter/cmd/ac
```

## Configuration

Before using the tool, you must configure an AI provider and provide the necessary API key.

1.  **API Key**: The tool requires an API key for the selected provider. For the Gemini provider, create a `.env` file in the project root and add your key:
    ```text
    GEMINI_API_KEY=your_api_key_here
    ```

2.  **Set AI Provider**: Use the `provider set` command to choose your AI provider. You will be prompted to select from a list of supported providers.

    ```bash
    # Set the AI provider via an interactive selector
    ac provider set
    ```

3.  **Verify Configuration**: You can check the currently configured provider at any time.

    ```bash
    # Get the currently configured provider
    ac provider get
    ```

## Usage

The typical workflow involves generating the project context first, followed by generating comments or a README file.

### 1. Generate Project Context

This command scans your Go project, analyzes the file structure and content, and saves a context summary that the AI will use in subsequent steps.

```bash
# Scan the project and generate context
ac context gen
```

### 2. Generate Code Comments

After the context is generated, you can use this command to automatically add comments to your Go source files. The tool identifies files that need comments and applies them.

```bash
# Generate comments for Go files
ac comments gen
```

### 3. Generate README File

This command uses the project context to generate a new `README.md` file for your project.

```bash
# Generate a README.md file
ac readme gen
```

## Project Structure

The project is organized into command-line definitions and internal packages that handle the core logic.

```text
.
├── cmd/
│   ├── ac/
│   │   └── main.go         # Main application entry point
│   ├── comments.go       # Defines the 'comments' command
│   ├── context.go        # Defines the 'context' command
│   ├── provider.go       # Defines the 'provider' command
│   ├── readme.go         # Defines the 'readme' command
│   └── root.go           # Root Cobra command setup
├── internal/
│   ├── ai/               # AI provider interfaces and implementations (Gemini)
│   ├── config/           # Application configuration management
│   ├── contextstore/     # Logic for saving and loading project context
│   ├── prompt/           # AI prompt templates and builders
│   ├── scanner/          # File system scanning and processing utilities
│   └── ui/               # User interface components like selectors
├── go.mod
└── LICENSE
```

## License

This project is available under the license specified in the `LICENSE` file.