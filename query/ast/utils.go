package ast

import "helldb/engine/types"

func ExtractToBaseType(expression ValueExpression) types.BaseType {
	var baseTypeElement types.BaseType
	switch expression.(type) {
	case *IntegerLiteral:
		baseTypeElement = expression.(*IntegerLiteral).ToBaseType()
	case *StringLiteral:
		baseTypeElement = expression.(*StringLiteral).ToBaseType()
	case *BooleanLiteral:
		baseTypeElement = expression.(*BooleanLiteral).ToBaseType()
	case *CollectionLiteral:
		baseTypeElement = expression.(*CollectionLiteral).ToBaseType()
	}
	return baseTypeElement
}
