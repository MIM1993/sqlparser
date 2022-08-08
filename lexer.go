/*
@Time : 2022/8/5 下午8:15
@Author : MuYiMing
@File : lexer
@Software: GoLand
*/
package sqlparser

type lexer struct {
	input string

	start    int
	position int
	width    int

	token chan lexToken
}
