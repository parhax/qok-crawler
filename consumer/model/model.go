package model

import (
	"database/sql"

	"qok.com/crawler/consumer/db"
)

type CrawlResult struct {
	Url   string
	Title string
}

//StoreInDB store crawling result into database
func (crawler *CrawlResult) StoreInDb() {
	msql, err := db.GetMysqlConnection()
	if err != nil {
		panic(err.Error())
	}
	defer msql.Close()

	checkForTableExistance(msql)

	var query = "INSERT IGNORE INTO title_crawling_results (`url`,`title`) VALUES (?, ?)"

	insert, err := msql.Query(query, crawler.Url, crawler.Title)

	if err != nil {
		panic(err.Error())
	}

	insert.Close()

}

func checkForTableExistance(msql *sql.DB) {
	var tableQuery = `CREATE TABLE IF NOT EXISTS title_crawling_results (
		id int(11) NOT NULL auto_increment,   
		url  varchar(100) NOT NULL default '',
		title varchar(20) NOT NULL default '',    
		 PRIMARY KEY  (id)
	  );`
	create, err := msql.Query(tableQuery)
	if err != nil {
		panic(err.Error())
	}
	create.Close()
}
