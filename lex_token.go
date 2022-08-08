/*
@Time : 2022/8/5 下午8:16
@Author : MuYiMing
@File : lex_token
@Software: GoLand
*/
package sqlparser

type lexToken struct {
	tokenType tokenType
	vlaue     string
}

type tokenType int

const (
	tokenError            tokenType = iota
	tokenSpace                      // whitespace
	tokenIdentifier                 // table or column name
	tokenEnd                        // the end of the input
	tokenEquals                     // "=="
	tokenAssign                     // "="
	tokenDelimeter                  // ','
	tokenLeftParenthesis            // '('
	tokenRightParenthesis           // ')'
	tokenInteger                    // integer
	tokenString                     // string including quotes
	tokenAnd                        // AND
	tokenInsert                     // INSERT
	tokenInto                       // INTO
	tokenSelect                     // SELECT
	tokenDelete                     // DELETE
	tokenFrom                       // FROM
	tokenWhere                      // WHERE
	tokenLimit                      // LIMIT
	tokenValues                     // VALUES
	tokenUpdate                     // UPDATE
	tokenSet                        // SET
	tokenCreate                     // CREATE
	tokenDrop                       // DROP
	tokenTable                      // TABLE
	tokenTypeInteger                // INTEGER
	tokenTypeString                 // STRING
	tokenEngine                     // ENGINE
	tokenLSM                        // LSM
	tokenBPTree                     // B+ tree
)
