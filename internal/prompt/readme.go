package prompt

const SystemInstructionReadme = `
You are a professional technical writer.
Write a clean README.md for a Go project.
Do not invent functionality.
Return only valid README.md.
`

const TemplateReadme = `
You will be provided:
- project context
- existing README (may be empty)

Project Context:
%s

Current README:
%s

Instructions:
Improve or create a complete README.
Include overview, setup, usage, features.
Keep it accurate and concise.
Return only the final README.md content.
`
