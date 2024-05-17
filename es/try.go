package es

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"strconv"
//
//	"github.com/olivere/elastic/v7"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//var ElasticClient *elastic.Client
//
//func init() {
//	var err error
//	ElasticClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9100"), elastic.SetSniff(false))
//	if err != nil {
//		log.Panicln("es连接失败" + err.Error())
//	}
//}
//
//// elasticSearch 添加数据
//func ElasticSearchAdd(index string, data interface{}) error {
//	_, err := ElasticClient.Index().Index(index).BodyJson(data).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// elasticSearch 查询数据
//func ElasticSearchSearch(index string, query elastic.Query) (*elastic.SearchResult, error) {
//	result, err := ElasticClient.Search().Index(index).Query(query).Do(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//// elasticSearch 高亮查询数据
//func ElasticSearchSearchWithHighlight(index string, query elastic.Query, highlight *elastic.Highlight) (*elastic.SearchResult, error) {
//	elastic.NewHighlight().Field("VideoName")
//	elastic.NewMatchQuery("VideoName", "")
//	result, err := ElasticClient.Search(index).Query(query).Highlight(highlight).Do(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//// sum 求和
//func Sum() {
//	sum := elastic.NewSumAggregation().Field("VideoWatch")
//	res, _ := ElasticClient.Search().Index("video").Size(0).Aggregation("one", sum).Do(context.Background())
//	ret, _ := ElasticClient.Count().Index("video").Query(elastic.NewMatchAllQuery()).Do(context.Background())
//	ref, _ := json.MarshalIndent(res, "", "")
//	rea, _ := json.MarshalIndent(ret, "", "")
//	all := make(map[string]interface{})
//	json.Unmarshal(ref, &all)
//	fmt.Println(res)
//	fmt.Println(ret)
//	fmt.Println(string(rea))
//	fmt.Println(all)
//}
//
//// elasticsearch ik分词器
//func ElasticSearchSearchWithIk(name string) {
//	ElasticClient.CreateIndex(name).Do(context.Background())
//	mapping := `{
//        "properties": {
//            "title": {
//                "type": "text",
//                "analyzer": "ik_max_word"
//            },
//            "sub_title": {
//                "type": "text",
//                "analyzer": "ik_max_word"
//            }
//        }
//    }`
//	res, err := ElasticClient.PutMapping().Index(name).BodyString(mapping).Do(context.Background())
//	//res, err := ElasticClient.PutMapping().Index(name).BodyJson(map[string]interface{}{
//	//	"settings": map[string]interface{}{
//	//		"analysis": map[string]interface{}{
//	//			"analyzer": map[string]interface{}{
//	//				"default": map[string]interface{}{
//	//					"analyzer": "ik_max_word",
//	//				},
//	//			},
//	//		},
//	//	},
//	//	"mappings": map[string]interface{}{
//	//		"properties": map[string]interface{}{
//	//			"title": map[string]interface{}{
//	//				"type":     "text",
//	//				"analyzer": "ik_max_word",
//	//			},
//	//		},
//	//	},
//	//}).Do(context.Background())
//	fmt.Println(err)
//	fmt.Println(res)
//}
//
//var db *gorm.DB
//
//func init() {
//	db, _ = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/zhuangaosi"), &gorm.Config{})
//}
//
//func main() {
//	//ElasticSearchSearchWithIk("video_a")
//	//List(1, 1, 1)
//	//for _, first := range firsts {
//	//
//	//}
//	//fmt.Println(firsts)
//	//b, _ := json.Marshal(firsts)
//	//var data []map[string]interface{}
//	//json.Unmarshal(b, &data)
//	//fmt.Println(string(b))
//	Sum()
//}
//
//func (Commen) TableName() string {
//	return "commen"
//}
//
//type Lists struct {
//	Main     Commen
//	Children []*Lists
//}
//
//type Commen struct {
//	Id        int
//	VideoId   int
//	EpisodeId int
//	Content   string
//	BelongId  int
//	Level     int
//}
//
//var firsts = []*Lists{}
//
//func List(videId int, episodeId int, level int) *Lists {
//	all := []Commen{}
//	db.Where("video_id = ? AND episode_id = ?", videId, episodeId).Find(&all)
//	for i := level - 1; i < len(all); i++ {
//		one := Commen{}
//		two := []Commen{}
//		db.Where("id = ?", all[i].BelongId).Find(&one)
//		db.Where("belong_id = ?", all[i].Id).Find(&two)
//		if one.Id != 0 && len(two) == 0 {
//			lf := &Lists{
//				Main:     all[i],
//				Children: make([]*Lists, 0),
//			}
//			return lf
//		} else if one.Id != 0 && len(two) != 0 && all[i].Level != 0 {
//			list := List(videId, episodeId, level+1)
//			second := NewList(all[i])
//			second.Children = append(second.Children, list)
//			return second
//		}
//		first := NewList(all[i])
//		firsts = append(firsts, first)
//		res := List(videId, episodeId, level+1)
//		first.Children = append(first.Children, res)
//	}
//	return nil
//}
//
//func NewList(commen Commen) *Lists {
//	return &Lists{
//		Main:     commen,
//		Children: make([]*Lists, 0),
//	}
//}
//
//func kk() {
//	var res *elastic.BulkService
//	var lists = []Lists{}
//	res = ElasticClient.Bulk()
//	for i, list := range lists {
//		request := elastic.NewBulkCreateRequest().Id(strconv.Itoa(i)).Doc(list)
//		res.Add(request)
//	}
//
//	res.Do(context.Background())
//}
