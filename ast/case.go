package ast

import "github.com/raviqqe/stg/types"

// Case is a case expression.
type Case interface {
	Expression
	Expression() Expression
	Type() types.Type
	DefaultAlternative() (DefaultAlternative, bool)
}
