package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{1, "Article 1", "Article 1 body"},
	article{2, "Article 2", "Article 2 body"},
	article{3, "ТЕСТИРУЕМ", "ВООООООООООООООООООООООООТ ЭЭЭЭЭЭЭЭЭЭЭЭЭЭЭТО НАСТОЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЩИЙ ТЕЕЕЕЕЕЕЕЕСТ"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
