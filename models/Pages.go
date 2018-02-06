package models

import "database/sql"

type Page struct {
	Title string
	Body  []byte
}

type H map[string]interface{}

type PageCollection struct {
	Pages []Page `json:"items"`
}

//func (p *Page) save(db *sql.DB) error  {
//	return PutPage(db, p.Title, p.Body)
//}
func (page *Page) Save(db *sql.DB) error {
	return PutPage(db, page.Title, page.Body)
}

func GetPage(db *sql.DB, title string) Page {
	query := "SELECT * FROM page WHERE title = ?"
	rows, err := db.Query(query, title)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	page := Page{}

	if rows.Next() {
		err2 := rows.Scan(&page.Title, &page.Body)
		if err2 != nil {
			panic(err2)
		}
		return page
	}

	return page
}

func GetPages(db *sql.DB) PageCollection {
	query := "SELECT * FROM page"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := PageCollection{}

	for rows.Next() {
		page := Page{}
		err2 := rows.Scan(&page.Title, &page.Body)

		if err2 != nil {
			panic(err2)
		}
		result.Pages = append(result.Pages, page)
	}

	return result
}

func PutPage(db *sql.DB, title string, body []byte) error {
	query := "INSERT INTO page(title, body) VALUES(?,?);"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(query)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	_, err2 := stmt.Exec(title, body)
	// Exit if we get an error
	if err2 != nil {
		println(title, body)
		return PostPage(db, title, body)
	}

	return nil
}
func PostPage(db *sql.DB, title string, body []byte) error {
	query := "UPDATE page SET body = ? WHERE title = ?;"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(query)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	_, err2 := stmt.Exec(body, title)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	p := GetPage(db, title)
	println(stmt)
	println(p.Body, p.Title)
	return nil
}
