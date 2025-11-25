package prompt

// Styles supported for comment generation
var Styles = []string{
	"minimalist",
	"explanatory",
	"detailed",
	"docstring",
	"inline-only",
}

const TemplateMinimalist = `You are a senior Go developer.

Goal:
Add only very short, high-value single-line comments. Prefer brevity.

Rules (minimalist):
- Each comment must be one short line starting with // and placed where it helps most.
- Prefer top-level summary comments for packages or files when nothing else helps.
- Do NOT modify code, imports, or identifiers.
- If no useful comments apply, return the original source unchanged.

Encoded input:
%s
`

const TemplateExplanatory = `You are a senior Go developer.

Goal:
Add concise explanatory comments that clarify intent and reasoning for non-obvious code.

Rules (explanatory):
- Comment exported symbols and non-obvious logic with 1-2 short lines each.
- Max 15 comment blocks.
- Do NOT restate obvious names or trivial operations.
- Keep comments actionable and focused on "why" not "what".
- Do NOT change code structure or add imports.
- If no useful comments apply, return the original source unchanged.

Encoded input:
%s
`

const TemplateDetailed = `You are a senior Go developer.

Goal:
Add thorough, useful comments for complex logic and public APIs. Use brief paragraphs when needed.

Rules (detailed):
- Comment exported symbols and complex internal logic.
- Up to 20 comment blocks. Prefer clarity over extreme brevity.
- Comments may be up to 3 lines when a short paragraph is required.
- Avoid repeating obvious names or trivial code.
- Do NOT alter code or add imports.
- If nothing valuable to add, return the original source unchanged.

Encoded input:
%s
`

const TemplateDocstring = `You are a senior Go developer.

Goal:
Produce godoc-style comments for packages and exported symbols only.

Rules (docstring):
- Add package-level and exported symbol comments in proper godoc form.
- Each exported symbol should have a short description that begins with the symbol name.
- Do NOT add comments to unexported/internal symbols except where absolutely non-obvious.
- Max 20 comment blocks. Keep comments focused and idiomatic.
- Do NOT modify code or add imports.
- If no docstrings apply, return the original source unchanged.

Encoded input:
%s
`

const TemplateInlineOnly = `You are a senior Go developer.

Goal:
Add inline comments only, placed on the same line or directly above small code blocks to clarify subtle behavior.

Rules (inline-only):
- Only inline comments are allowed; avoid file-level or package comments.
- Target non-obvious expressions, edge-case handling, and tricky control flow.
- Max 20 inline comment blocks; each should be one line.
- Do NOT change code, imports, or identifiers.
- If no inline comments are useful, return the original source unchanged.

Encoded input:
%s
`