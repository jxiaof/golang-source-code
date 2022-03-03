package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	maxLimit    int32 = 2
	FieldsOrder int64 = 1
)

var (
	ArgoStatus   = "argoStatus"
	ArgoLog      = "argoLog"
	DBAtmodels   = "atmodels-test"
	CollAtmodels = "atmodels"
	CollProjs    = "projs"
	CollMrpVers  = "modelrp_versions"
	CollMrpHis   = "testhistorys"
)

type Scoll struct {
	Cli    *mongo.Client
	ColMap map[string]*mongo.Collection
}

type SListParams struct {
	Skip   int64
	Limit  int64
	Sorts  string
	Fields string
}

type SMongoListParams struct {
	Filter  string
	Options SListParams
}

type SListRes struct {
	Total int64
	Data  interface{}
}

type SfindRes struct {
	Count int64
	Data  []primitive.M
}
type SGetParams struct {
	Skip   int64
	Sorts  string
	Fields string
}

type SMongoGetParams struct {
	Filter  string
	Options SGetParams
}

type SUpdateParams struct {
	Multi  bool
	Upsert bool
}

type SMongoUpdateParams struct {
	Filter  string
	Update  string
	Options SUpdateParams
}

type SDeleteParams struct {
	Multi bool
}

type SMongoDeleteParams struct {
	Filter  string
	Options SUpdateParams
}

type MogCrud interface {
	CheckExist(colName, filter string) (bool, error)
	List(colName string, p *SMongoListParams) (*SfindRes, error)
	Get(colName string, p *SMongoGetParams) (*primitive.M, error)
	Create(colName string, doc []interface{}) ([]interface{}, error)
	Update(colName string, p *SMongoUpdateParams) (*mongo.UpdateResult, error)
	Delete(colName string, p *SMongoDeleteParams) (*mongo.DeleteResult, error)
}

// 打印集合的信息
func (coll *Scoll) Desc() (interface{}, error) {
	res := make(map[string]string, 0)
	// collection := coll.ColMap[colName]
	var indexView *mongo.IndexView

	// Specify the MaxTime option to limit the amount of time the operation can
	// run on the server
	opts := options.ListIndexes().SetMaxTime(2 * time.Second)
	cursor, err := indexView.List(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// Get a slice of all indexes returned and print them out.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println(results)
	res["success"] = "true"
	// res["results"] = results
	return res, nil
}

// 删除一个document
func (coll *Scoll) Delete(colName string, p *SMongoDeleteParams) (*mongo.DeleteResult, error) {
	var err error
	var res *mongo.DeleteResult
	collection := coll.ColMap[colName]
	if p.Filter == "" {
		return nil, fmt.Errorf("filter is empty")
	}
	// 解析过滤条件
	fil, err := Json2Bson(p.Filter)
	if err != nil {
		err = fmt.Errorf("parser filter json to bson error: %v", err)
		return nil, err
	}
	opts := options.DeleteOptions{}
	if p.Options.Multi {
		res, err = collection.DeleteMany(context.TODO(), fil, &opts)
	} else {
		res, err = collection.DeleteOne(context.TODO(), fil, &opts)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (coll *Scoll) Update(colName string, p *SMongoUpdateParams) (*mongo.UpdateResult, error) {
	var err error
	var res *mongo.UpdateResult
	collection := coll.ColMap[colName]
	if p.Filter == "" {
		return nil, fmt.Errorf("filter is empty")
	}
	if p.Update == "" {
		return nil, fmt.Errorf("update is empty")
	}
	// 解析过滤条件
	fil, err := Json2Bson(p.Filter)
	if err != nil {
		err = fmt.Errorf("parser filter json to bson error: %v", err)
		return nil, err
	}
	// 解析更新条件
	upd, err := Json2Bson(p.Update)
	if err != nil {
		err = fmt.Errorf("parser update json to bson error: %v", err)
		return nil, err
	}
	opts := options.UpdateOptions{}

	if p.Options.Upsert {
		opts.SetUpsert(true)
	}
	if p.Options.Multi {
		res, err = collection.UpdateMany(context.TODO(), fil, upd, &opts)
	} else {
		res, err = collection.UpdateOne(context.TODO(), fil, upd, &opts)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 提供了一个简单的操作mongodb的接口,检查document是否存在
func (coll *Scoll) CheckExist(colName, filter string) (bool, error) {
	var err error
	var res bson.M
	opts := options.FindOneOptions{}
	fil, err := Json2Bson(filter)
	if err != nil {
		return false, err
	}
	err = coll.ColMap[colName].FindOne(context.TODO(), fil, &opts).Decode(&res)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (coll *Scoll) Get(colName string, p *SMongoGetParams) (*primitive.M, error) {
	var err error
	var res bson.M

	if p.Filter == "" {
		return nil, fmt.Errorf("filter is empty")
	}
	opts := options.FindOneOptions{}
	if p.Options.Skip > 0 {
		opts.Skip = &p.Options.Skip
	}
	if p.Options.Sorts != "" {
		sort, err := Json2Bson(p.Options.Sorts)
		if err != nil {
			err = fmt.Errorf("parser sorts json to bson error: %v", err)
			return nil, err
		}
		opts.Sort = sort
	}
	if p.Options.Fields != "" {
		fields, err := Json2Bson(p.Options.Fields)
		if err != nil {
			err = fmt.Errorf("parser fields json to bson error: %v", err)
			return nil, err
		}
		opts.Projection = fields
	}
	// 解析过滤条件
	fil, err := Json2Bson(p.Filter)
	if err != nil {
		err = fmt.Errorf("json to bson error: %v", err)
		return nil, err
	}
	collection := coll.ColMap[colName]
	err = collection.FindOne(context.TODO(), fil, &opts).Decode(&res)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, nil
	}
	return &res, nil
}

// 创建一个document
func (coll *Scoll) Create(colName string, doc []interface{}) ([]interface{}, error) {
	collection := coll.ColMap[colName]
	if len(doc) == 1 {
		res, err := collection.InsertOne(context.TODO(), doc[0])
		if err != nil {
			return nil, err
		}
		if res.InsertedID == nil {
			return nil, fmt.Errorf("return insert id is empty")
		}

		fmt.Printf("inserted one document with ID %v\n", res.InsertedID)
		// return &res, nil
		return []interface{}{res.InsertedID}, nil
	} else {
		res, err := collection.InsertMany(context.TODO(), doc)
		if err != nil {
			return nil, err
		}
		if res.InsertedIDs == nil {
			return nil, fmt.Errorf("return insert id is empty")
		}
		fmt.Printf("inserted many documents with IDs %v\n", res.InsertedIDs)
		// return &res.InsertedIDs, nil
		return res.InsertedIDs, nil
	}
}

// 提供一个列表接口
func (coll *Scoll) List(colName string, p *SMongoListParams) (*SfindRes, error) {
	var err error
	var data []bson.M

	if p.Filter == "" {
		return nil, fmt.Errorf("filter is empty")
	}
	opts := options.FindOptions{}
	if p.Options.Skip > 0 {
		opts.Skip = &p.Options.Skip
	}
	if p.Options.Limit > 0 {
		opts.Limit = &p.Options.Limit
	}
	if p.Options.Sorts != "" {
		sort, err := Json2Bson(p.Options.Sorts)
		if err != nil {
			return nil, err
		}
		fmt.Println("sort:", sort)
		opts.Sort = sort
	}
	if p.Options.Fields != "" {
		fields, err := Json2Bson(p.Options.Fields)
		if err != nil {
			return nil, err
		}
		fmt.Println("fields:", fields)
		opts.Projection = fields
	}

	fil, err := Json2Bson(p.Filter)
	if err != nil {
		err = fmt.Errorf("json to bson error: %v", err)
		return nil, err
	}
	collection := coll.ColMap[colName]
	cur, err := collection.Find(context.TODO(), &fil, &opts)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, nil
	}
	defer cur.Close(context.TODO())
	if err := cur.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	// 获取数据总数
	count, err := collection.CountDocuments(context.Background(), &fil)
	if err != nil {
		return nil, err
	}

	result := SfindRes{
		Count: count,
		Data:  data,
	}
	return &result, nil
}
func Json2Bson(s string) (interface{}, error) {
	/*
		将json字符串转换为bson
	*/
	var doc interface{}
	err := bson.UnmarshalExtJSON([]byte(s), true, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetCli() (*mongo.Client, error) {
	var err error
	var cli *mongo.Client
	cli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:a6oNBpPs@10.13.129.218:27017"))
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func GetColl() (*Scoll, error) {
	var err error
	var scol Scoll
	scol.ColMap = make(map[string]*mongo.Collection)
	scol.Cli, err = GetCli()
	if err != nil {
		return nil, err
	}
	scol.ColMap["atmodels"] = scol.Cli.Database(DBAtmodels).Collection(CollAtmodels)
	scol.ColMap["projs"] = scol.Cli.Database(DBAtmodels).Collection(CollProjs)
	scol.ColMap["vers"] = scol.Cli.Database(DBAtmodels).Collection(CollMrpVers)
	scol.ColMap["molhis"] = scol.Cli.Database(DBAtmodels).Collection(CollMrpHis)
	scol.ColMap["argoLog"] = scol.Cli.Database(DBAtmodels).Collection(ArgoLog)
	scol.ColMap["argoStatus"] = scol.Cli.Database(DBAtmodels).Collection(ArgoStatus)
	return &scol, nil
}

// NewIMogDao 初始化
func NewMogCrud() *Scoll {
	coll, err := GetColl()
	if err != nil {
		panic(err)
	}
	return coll
}

func testCheckExist() {
	var err error
	var fil string
	var dao MogCrud
	dao = NewMogCrud()

	fil = `{"_id": "resnet"}`
	flg, err := dao.CheckExist("projs", fil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("+++++++", flg)
}

func testList() {
	var err error
	var p SMongoListParams
	var dao MogCrud
	dao = NewMogCrud()
	p.Filter = fmt.Sprintf(`{"_id": {"$regex": "%s"}}`, "r")
	p.Options.Sorts = `{"create_time": -1, "_id": -1}`
	p.Options.Limit = 10
	p.Options.Skip = 0
	p.Options.Fields = `{"_id": 1, "name": 1}`
	res, err := dao.List("projs", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------", res.Count, res.Data)
}

func testGet() {
	var p SMongoGetParams

	var dao MogCrud
	dao = NewMogCrud()
	p.Filter = fmt.Sprintf(`{"_id": {"$regex": "%s"}}`, "r")
	p.Options.Fields = `{"_id": 1, "name": 1}`
	p.Options.Skip = 0
	p.Options.Sorts = `{"create_time": -1, "_id": -1}`
	res, err := dao.Get("projs", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("++++++++++++", *res)
}

func testCreate() {
	var dao MogCrud
	dao = NewMogCrud()
	doc0 := make(map[string]interface{})
	doc := make(map[string]interface{})
	doc2 := make(map[string]interface{})
	doc3 := make(map[string]interface{})
	doc["_id"] = "111111new"
	doc2["_id"] = "222222new"
	doc3["_id"] = "333333new"
	doc["name"] = "1"
	doc["history"] = []string{"1", "2"}
	d := []interface{}{&doc, &doc2, &doc3}
	r, err := dao.Create("projs", d)
	if err != nil {
		fmt.Println(err)
	}
	doc0["_id"] = "000000new"
	r2, err := dao.Create("projs", []interface{}{&doc0})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("++++++++++++", r, r2)
}

func testUpdate() {
	var p SMongoUpdateParams
	var dao MogCrud
	dao = NewMogCrud()
	p.Filter = fmt.Sprintf(`{"_id": {"$regex": "%s"}}`, "new")
	p.Update = `{"$set": {"name": "++++++1111111", "history": ["1", "2", "3"]}}`
	p.Options.Multi = true
	p.Options.Upsert = false
	r, err := dao.Update("projs", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("++++++++++++", r)
}

func testDelete() {
	var p SMongoDeleteParams
	var dao MogCrud
	dao = NewMogCrud()
	p.Filter = fmt.Sprintf(`{"_id": {"$regex": "%s"}}`, "rs")
	p.Options.Multi = true
	r, err := dao.Delete("projs", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("++++++++++++", r)
}

func testDesc() {
	var err error
	coll := NewMogCrud()
	res, err := coll.Desc()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("++++++++++++", res)
}
func main() {
	// testCheckExist()
	// testList()
	// testGet()
	// testCreate()
	// testUpdate()
	// testDelete()
	testDesc()

}
