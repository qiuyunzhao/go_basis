package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"go_basis/30_数据库操作/4_influxDB/influxUtils"
	"math/rand"
	"time"
)

func main() {

	// ==================================================== 获取数据库连接 ==============================================
	con := influxUtils.Conn{
		InfluxAddr: "http://127.0.0.1:8086",
	}
	influxClient := con.ConnInflux()

	defer influxClient.Close()

	// ===================================================== 执行SQL ===================================================
	// SQL语句语法见官网 ：https://docs.influxdata.com/influxdb/v1.7/query_language/data_download/

	//数据库操作
	//sql := `SHOW databases`
	//sql := `CREATE DATABASE NOAA_water_database WITH DURATION 5d REPLICATION 2 SHARD DURATION 1h NAME liquid`
	//sql := `DROP DATABASE NOAA_water_database`

	//数据表操作
	//sql := `SHOW measurements ON Weather`

	//存储策略
	//sql := `SHOW RETENTION POLICIES`

	//连续查询
	//sql := `SHOW CONTINUOUS QUERIES`
	//sql := `DROP CONTINUOUS QUERY "cq_basic" ON "Weather"`
	sql := `CREATE CONTINUOUS QUERY "cq_basic" ON "Weather"
    BEGIN
        SELECT mean("温度") AS "mean_温度" INTO "Weather"."12hours"."江苏" FROM "Weather"."1hour"."江苏" WHERE "城市"='无锡' GROUP BY time(30s)
    END`

	//查询数据
	//sql := `SELECT mean("风速") AS "mean_风速" FROM "Weather"."1hour"."江苏" WHERE time > now()-9s AND "城市"='无锡' GROUP BY time(3s) FILL(null)`
	//sql := `SELECT *  FROM "Weather"."1hour"."江苏" WHERE time > now()-10s`
	//sql := `SELECT COUNT("温度") FROM "Weather"."1hour"."江苏"`
	//sql := ` SELECT "温度" FROM "江苏" LIMIT 2`

	result, err := influxUtils.ExecuteSQL(influxClient, sql) //没有获取到复核的数据结果：[{[] [] }]
	if err == nil {
		fmt.Println(result)
	}

	// ==================================================== 插入数据 ====================================================
	batchPoint := influxUtils.BatchPoints{
		Database:        "Weather",
		Precision:       "s",
		RetentionPolicy: "1hour",
		//WriteConsistency:,
		Measurement: "山东",
	}
	go insertPoint(batchPoint, influxClient)
	batchPoint.Measurement = "江苏"
	go insertPoint1(batchPoint, influxClient)
	select {}
}

//======================================================================================================================
func insertPoint1(batchPoint influxUtils.BatchPoints, client client.Client) {
	for {
		time.Sleep(time.Millisecond * 1000)

		rand.Seed(time.Now().UnixNano())
		fields := map[string]interface{}{
			"天气": "小雨",
			"温度": rand.Float64()*30 + 2,
			"风速": rand.Float64() * 10,
		}
		tags := map[string]string{"城市": "苏州"}
		dataPoint := influxUtils.DataPoint{
			Tags:      tags,
			Fields:    fields,
			Timestamp: time.Now(),
		}
		batchPoint.Point = dataPoint

		//核心
		batchPoint.WritesPoints(client)
		//fmt.Println(batchPoint.Measurement, "---", dataPoint.Tags, "--", dataPoint.Fields, "--", dataPoint.Timestamp)

		//------------------------------------------------------------------------------------------
		rand.Seed(time.Now().UnixNano())
		fields1 := map[string]interface{}{
			"天气": "晴转阴",
			"温度": rand.Float64()*30 - 5,
			"风速": rand.Float64() * 10,
		}
		tags1 := map[string]string{"城市": "无锡"}
		dataPoint1 := influxUtils.DataPoint{
			Tags:      tags1,
			Fields:    fields1,
			Timestamp: time.Now(),
		}
		batchPoint.Point = dataPoint1

		batchPoint.WritesPoints(client)
		//fmt.Println(batchPoint.Measurement, "---", dataPoint1.Tags, "--", dataPoint1.Fields, "--", dataPoint1.Timestamp)
	}
}

func insertPoint(batchPoint influxUtils.BatchPoints, client client.Client) {
	for {
		time.Sleep(time.Millisecond * 1000)

		rand.Seed(time.Now().UnixNano())
		fields := map[string]interface{}{
			"天气": "晴",
			"温度": rand.Float64()*40 - 10,
			"风速": rand.Float64() * 20,
		}
		tags := map[string]string{"城市": "济南"}
		dataPoint := influxUtils.DataPoint{
			Tags:      tags,
			Fields:    fields,
			Timestamp: time.Now(),
		}
		batchPoint.Point = dataPoint

		//核心
		batchPoint.WritesPoints(client)
		//fmt.Println(batchPoint.Measurement, "---", dataPoint.Tags, "--", dataPoint.Fields, "--", dataPoint.Timestamp)

		//------------------------------------------------------------------------------------------
		rand.Seed(time.Now().UnixNano())
		fields1 := map[string]interface{}{
			"天气": "阴",
			"温度": rand.Float64()*40 - 10,
			"风速": rand.Float64() * 30,
		}
		tags1 := map[string]string{"城市": "青岛"}
		dataPoint1 := influxUtils.DataPoint{
			Tags:      tags1,
			Fields:    fields1,
			Timestamp: time.Now(),
		}
		batchPoint.Point = dataPoint1

		batchPoint.WritesPoints(client)
		//fmt.Println(batchPoint.Measurement, "---", dataPoint1.Tags, "--", dataPoint1.Fields, "--", dataPoint1.Timestamp)
	}
}
