package model

import (
	"database/sql"

	"qok.com/crawler/consumer/db"
)

type Crawler struct {
	url   string
	title string
}

func (crawler *Crawler) SetUrl(str string) {
	crawler.url = str
}
func (crawler *Crawler) SetTitle(str string) {
	crawler.title = str
}

// func (crawler *Crawler) DoCrawl(sUrl string) (string, error) {
// 	resp, err := http.Get(sUrl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	strBody := string(body)
// 	// re := regexp.MustCompile("<title>(.*?)</title>")
// 	re := regexp.MustCompile("<title*>(.*?)</title>")
// 	match := re.FindStringSubmatch(strBody)
// 	if len(match) <= 0 {
// 		// fmt.Printf("Could not find any title for %s ", sUrl)
// 		return Error("Could not find any title for %s")
// 	} else {
// 		// fmt.Printf("%s ", match[0])
// 		return match[0], nil
// 	}

// }
func (crawler *Crawler) StoreInDb() {
	msql, err := db.GetMysqlConnection()
	if err != nil {
		panic(err.Error())
	}
	defer msql.Close()

	checkForTableExistance(msql)

	var query = "INSERT IGNORE INTO title_crawling_results (`url`,`title`) VALUES (?, ?)"

	insert, err := msql.Query(query, crawler.url, crawler.title)

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
