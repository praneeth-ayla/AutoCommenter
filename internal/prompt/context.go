package prompt

const TemplateFileContext = `
Analyze the Go file and return a single JSON object:
- path: string
- file_name: string
- summary: really small summary about file purpose
- exports: list of exported symbols
- imports: list of imported package paths

Do not include any fields not listed.
Do not include explanations.

Path:
%s

Content:
%s
`