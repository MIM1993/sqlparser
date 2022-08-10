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

var tokenToString = map[tokenType]string{
	tokenError:            "error",
	tokenSpace:            "spqce",
	tokenIdentifier:       "identifier",
	tokenEnd:              "end",
	tokenEquals:           "equals",
	tokenAssign:           "assign",
	tokenDelimeter:        "delimeter",
	tokenLeftParenthesis:  "leftParenthesis",
	tokenRightParenthesis: "rightParenthesis",
	tokenInteger:          "integer",
	tokenString:           "string",
	tokenAnd:              "AND",
	tokenInsert:           "INSERT",
	tokenInto:             "INTO",
	tokenSelect:           "SELECT",
	tokenDelete:           "DELETD",
	tokenFrom:             "FROM",
	tokenWhere:            "WHERE",
	tokenLimit:            "LIMIT",
	tokenValues:           "VALUES",
	tokenUpdate:           "UPDATE",
	tokenSet:              "SET",
	tokenCreate:           "CREATE",
	tokenDrop:             "DROP",
	tokenTable:            "TABLE",
	tokenTypeInteger:      "typeInteger",
	tokenTypeString:       "typeString",
	tokenEngine:           "ENGINE",
	tokenLSM:              "LSM",
	tokenBPTree:           "BPTree",
}

func (t tokenType) String() string {
	s, defined := tokenToString[t]
	if defined {
		return s
	}
	return "unknown"
}

const end = -1

/*
	tokenAnd:              "AND",
	tokenInsert:           "INSERT",
	tokenInto:             "INTO",
	tokenSelect:           "SELECT",
	tokenDelete:           "DELETD",
	tokenFrom:             "FROM",
	tokenWhere:            "WHERE",
	tokenLimit:            "LIMIT",
	tokenValues:           "VALUES",
	tokenUpdate:           "UPDATE",
	tokenSet:              "SET",
	tokenCreate:           "CREATE",
	tokenDrop:             "DROP",
	tokenTable:            "TABLE",
	tokenTypeInteger:      "typeInteger",
	tokenTypeString:       "typeString",
	tokenEngine:           "ENGINE",
	tokenLSM:              "LSM",
	tokenBPTree:           "BPTree",
*/

var keyWords = map[string]tokenType{
	"": tokenAnd,
}
