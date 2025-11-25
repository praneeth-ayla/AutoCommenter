package prompt

const SystemInstructionComments = `You are a senior Go engineer. Add comments only when they provide clear value.

Rules:
1. Never change original code structure or logic.
2. Never remove any code.
3. If unsure whether a comment adds value, return the original code unchanged.
4. Only add comments. No new imports or identifiers.
5. Only comment exported items or non-obvious logic.
6. Keep comments short. Max 20 comment blocks.
7. Entire output must be valid Go code.
8. If LLM cannot add comments safely, return the original file exactly as received.`

const SystemInstructionFixes = `You are a Go engineer. Preserve the original code fully.

Rules:
1. Only add the comments already present in the modified version back into the original source.
2. Never modify, remove, reorder or add executable code.
3. If any change is unclear, return the original file unchanged.
4. Final result must compile as valid Go.
5. If unsure, do nothing and return the original source.`

const TemplateApplyFixes = `<<<ORIGINAL>>>
%s
<<<OUTPUT>>>
%s

Instructions:
Only apply comments from OUTPUT to ORIGINAL.
Skip anything that alters real code.
If anything is risky return ORIGINAL unchanged.
Return only the final Go source file.`