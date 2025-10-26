

# 🧠 AutoCommenter – AI-Powered Code Comment & Documentation Generator

### Overview

**AutoCommenter** is a command-line tool built in **Go** that automatically adds production-level comments and generates detailed documentation for any codebase.
It integrates with **Git** to identify updated files, analyzes their logic using an **AI model (GPT or local LLM)**, and inserts meaningful comments to improve code readability and maintainability.

The tool also supports automatic **Markdown documentation generation**, helping teams quickly create technical docs based on the actual code context.

---

### 🚀 Features

* **Automated Commenting**
  Scans the project and intelligently inserts detailed, context-aware comments in source files.

* **Smart Git Integration**
  Detects changed or updated files using `git status` and limits operations to only those files.

* **Documentation Generator**
  Creates clean, structured Markdown docs summarizing each function, class, and file.

* **Interactive Mode**
  If the model lacks context, it can request clarification or open specific files for user input.

* **License Assistant**
  Helps users choose a suitable open-source license and explains each option in simple terms.

* **Configurable Models**
  Works with GPT API or local LLMs (like Code LLaMA or GPT4All) for offline use.

---

### 🧩 Tech Stack

* **Language:** Go
* **Libraries:** Cobra (CLI), os/exec, io/fs
* **AI Integration:** OpenAI API / Local LLM (via Ollama or llama.cpp)
* **Version Control:** Git
* **Documentation Output:** Markdown

---

### ⚙️ How It Works

1. **Scan Repository**
   Walks through the project structure and identifies code files.

2. **Find Updated Files**
   Uses `git status --porcelain` to detect files modified since last commit.

3. **Generate Comments**
   Sends file content to the AI model for understanding and comment generation.

4. **Write Back**
   Inserts structured comments directly into code or outputs them to a preview file.

5. **Generate Documentation (Optional)**
   Creates Markdown summaries of key functions and logic flow across files.

---

### 🧰 Example Commands

```bash
# Add comments to all updated files
go run main.go --action=comment

# Generate documentation for the entire project
go run main.go --action=docs

# Use a specific model (e.g., local Code LLaMA)
go run main.go --action=comment --model=local
```

---

### 🧠 Example Workflow

```bash
> git status
M main.go
M routes/handlers.go

> go run main.go --action=comment
✅ Added detailed comments to 2 files

> go run main.go --action=docs
📄 Generated documentation at /docs/project_summary.md
```

---

### 🧱 Project Structure

```
AutoCommenter/
│
├── main.go              # Entry point for CLI
├── cmd/                 # CLI command definitions
├── internal/
│   ├── gitutils/        # Functions for Git integration
│   ├── fileparser/      # Reads and parses code files
│   ├── ai/              # LLM interaction logic
│   ├── docgen/          # Markdown documentation generator
│   └── utils/           # Helper utilities
├── go.mod
└── README.md
```

---

### 🌐 Future Enhancements

* Fine-tune LLMs for commenting style per language.
* Add project-level context caching for faster commenting.
* Enable real-time chat mode for code explanation & editing.
* VSCode extension for inline AI comments.

---

### 📜 License

AutoCommenter supports multiple open-source licenses.
You can generate or change the license interactively:

```bash
go run main.go --action=license
```

If you’re unsure which license fits your project, the AI can explain each option in simple terms.

