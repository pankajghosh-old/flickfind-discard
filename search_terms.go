package main

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	// "github.com/bmizerany/pq"
	// "github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	// "os"
	// "strings"
	// "github.com/jmoiron/sqlx"
)

type SearchTerm struct {
	id   int
	term string
}

// type SearchTerms struct {
// 	searchTerms []SearchTerm
// }

func (st *SearchTerm) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = st.id
	res["term"] = st.term
	return res
}

func GetAllSearchTerms() (terms []SearchTerm) {
	rows, err := db.Query("SELECT * FROM SearchTerms")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r_search_term SearchTerm
		err := rows.Scan(&r_search_term.id, &r_search_term.term)
		if err != nil {
			fmt.Println("Scan: %v", err)
		}
		terms = append(terms, r_search_term)
	}
	return terms
}

func AddSearchTerm(new_term string) {
	sql := fmt.Sprintf("INSERT INTO SearchTerms (term) SELECT '%s'"+
		"WHERE NOT EXISTS ( SELECT id FROM SearchTerms WHERE term = '%s');", new_term, new_term)

	fmt.Println("executing sql", sql)

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

}
