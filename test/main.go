package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Annotation ...
type Annotation struct {
	Status int
	File   string
	Text   string
	Gender int
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	findOptions := options.Find()
	// findOptions.SetLimit(2)
	collection := client.Database("maybe-data").Collection("annotations")
	var results []Annotation
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("delete.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Annotation
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		if elem.Status == -1 {
			results = append(results, elem)
			lineStr := fmt.Sprintf("%s", elem.File)
			fmt.Fprintln(w, lineStr)
		}
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())

	w.Flush()
	fmt.Printf("Found multiple documents (array of pointers): %#v\n", len(results))
}
