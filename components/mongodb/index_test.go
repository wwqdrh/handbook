package mongodb

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func TestModel(t *testing.T) {
	InitConnect()
	defer DisConnect()
	// 单个插入
	ash := Trainer{"Ash", 10, "Pallet Town"}
	InsertOneResult := NewMgo().InsertOne(ash)
	fmt.Println("Inserted a single document: ", InsertOneResult)

	// 插入多个值
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	trainers := []interface{}{misty, brock}
	insertManyResult := NewMgo().InsertMany(trainers)
	fmt.Println("Inserted multiple documents: ", insertManyResult)

	// 更新
	update := bson.D{
		{"$inc", bson.D{
			{"age", 999},
		}},
	}
	updateResult := NewMgo().UpdateOne("name", "Ash", update)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// 查询一个
	var result Trainer
	NewMgo().FindOne("name", "Ash").Decode(&result)
	fmt.Printf("Found a single document: %+v\n", result)

	// 查询总数
	name, size := NewMgo().Count()
	fmt.Printf(" documents name: %+v documents size %d \n", name, size)

	// 查询多个记录
	var results []*Trainer
	cur := NewMgo().FindAll(0, size, 1)
	defer cur.Close(context.TODO())
	if cur != nil {
		fmt.Println("FindAll err:", cur)
	}
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 遍历结果
	for k, v := range results {

		fmt.Printf("Found  documents  %d  %v \n", k, v)
	}

	// 删除文件
	deleteResult := NewMgo().DeleteMany("name", "Ash")
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult)
}
