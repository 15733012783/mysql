package es

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"time"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL("http://120.27.208.86:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}

func CreateBook(e1 interface{}, id string) {
	_, err := client.Index().
		Index("book").
		Id(id).
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
}

type Book struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ID        int         `json:"ID"`
				CreatedAt time.Time   `json:"CreatedAt"`
				UpdatedAt time.Time   `json:"UpdatedAt"`
				DeletedAt interface{} `json:"DeletedAt"`
				UserId    int         `json:"UserId"`
				BookName  string      `json:"BookName"`
				BookStars string      `json:"BookStars"`
				BookClass string      `json:"BookClass"`
				ClassID   int         `json:"ClassID"`
				Comment   int         `json:"Comment"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func FcSearch(text string) (Book, error) {
	serrch, err := client.Search("book").Query(elastic.NewMatchQuery("BookName", text)).Do(context.Background())
	if err != nil {
		return Book{}, err
	}
	indent, err := json.MarshalIndent(serrch, "", "")
	if err != nil {
		return Book{}, err
	}
	var book Book
	err = json.Unmarshal(indent, &book)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

type High struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ID        int         `json:"ID"`
				CreatedAt time.Time   `json:"CreatedAt"`
				UpdatedAt time.Time   `json:"UpdatedAt"`
				DeletedAt interface{} `json:"DeletedAt"`
				UserId    int         `json:"UserId"`
				BookName  string      `json:"BookName"`
				BookStars string      `json:"BookStars"`
				BookClass string      `json:"BookClass"`
				ClassID   int         `json:"ClassID"`
				Comment   int         `json:"Comment"`
			} `json:"_source"`
			Highlight struct {
				BookName []string `json:"BookName"`
			} `json:"highlight"`
		} `json:"hits"`
	} `json:"hits"`
}

func HSearch(text string) (High, error) {
	serrch, err := client.Search("book").
		Query(elastic.NewMatchQuery("BookName", text)).
		Highlight(elastic.NewHighlight().
			Field("BookName").
			PreTags("<span style='color:red'>").
			PostTags("</span>")).
		Do(context.Background())

	if err != nil {
		return High{}, err
	}
	indent, err := json.MarshalIndent(serrch, "", "")
	if err != nil {
		return High{}, err
	}
	var high High
	err = json.Unmarshal(indent, &high)
	if err != nil {
		return High{}, err
	}
	return high, nil
}
