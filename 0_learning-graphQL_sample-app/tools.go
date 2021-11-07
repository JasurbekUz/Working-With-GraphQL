//go:build tools
// +build tools

package tools

import _ "github.com/99designs/gqlgen"
// printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

// bu aslida go bilan ishkayigan app
// bizga kerakli arsalarni o'zi ishlatadi 