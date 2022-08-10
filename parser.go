/*
@Time : 2022/8/5 下午8:15
@Author : MuYiMing
@File : parser
@Software: GoLand
*/
package sqlparser

type parser struct {
	lexer *lexer

	//statement Statement

	err error
}


