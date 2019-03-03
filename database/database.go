package database

import (
	"database/sql"
	"github.com/andrzejd-pl/SimpleRestBlogBackend/usage"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Category struct {
	CategoryId          int    `column:"CategoryID";json:"category-id"`
	CategoryName        string `column:"CategoryName";json:"category-name"`
	CategoryDescription string `column:"CategoryDescription";json:"category-description"`
}

type Article struct {
	ArticleID      int      `column:"ArticleID";json:"article-id"`
	Category       Category `column:"CategoryID";json:"category"`
	ArticleTitle   string   `column:"ArticleTitle";json:"article-title"`
	ArticleContent string   `column:"ArticleContent";json:"article-content"`
}

func GetAllCategories() []Category {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name)
	defer db.Close()
	usage.CheckErr(err)

	rows, err := db.Query("select * from categories")
	usage.CheckErr(err)

	var categories []Category

	for rows.Next() {
		var model Category
		err := rows.Scan(&model.CategoryId, &model.CategoryName, &model.CategoryDescription)
		usage.CheckErr(err)

		categories = append(categories, model)
	}

	return categories
}

func GetAllArticles() []Article {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name)
	defer db.Close()
	usage.CheckErr(err)

	rows, err := db.Query("select * from articles")
	usage.CheckErr(err)

	var articles []Article

	for rows.Next() {
		var model Article
		var categoryId string
		err := rows.Scan(&model.ArticleID, &categoryId, &model.ArticleTitle, &model.ArticleContent)
		usage.CheckErr(err)
		categoryRows, err := db.Query("select * from categories where CategoryID = ?", categoryId)
		usage.CheckErr(err)

		if categoryRows.Next() {
			var category Category
			err := categoryRows.Scan(&category.CategoryId, &category.CategoryName, &category.CategoryDescription)
			usage.CheckErr(err)
			model.Category = category
		}

		articles = append(articles, model)
	}

	return articles
}
