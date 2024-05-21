# Static-Site-Generator
## Purpose
This project intends to take a series of markdown files and convert them to a series of html files that are ready to be hosted.

## Transpiler approach
1. Lexer converts file to series of tokens
1. Tokens are passed to Parser (will occur concurrently eventually)
1. Parser outputs new file

## Lexer approach
There are two possible ways to perform this:
- Use lexer with a single pass to grab each token
- Have lexer perform a second pass when inline values are detected

### Single pass 
- Only one pass meaning no reason to parse anything twice
- Can constatly output tokens(even though check would most likely be small)
- AST's typically do not care about parenthesis (Would be difficult)

### Multi pass
- Second pass would be small and can be a simple flag during initial parsing
- Most likely only minor slowdown (only slows down on a "large" token and only impacts that token(not line))
- More form fitting with most AST implementations

Due to the reasons discussed above, it makes more sense to go with a multipass approach

### Lexer Algo
1. Line is passed in
1. Begin parsed character by character
1. Make tokens as needed(appending to slice)
1. Return token list OR error(specifying line and value that broke it)
1. Eventually consider streaming tokens???

