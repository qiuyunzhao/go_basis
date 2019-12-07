package influxUtils

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

//===================================================== 获取连接 ======================================================
//数据库连接结构体
type Conn struct {
	InfluxAddr string //数据库连接地址 "http://127.0.0.1:8086"
	Username   string //用户名
	Password   string //密码
}

//连接数据库
func (conn *Conn) ConnInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     conn.InfluxAddr,
		Username: conn.Username,
		Password: conn.Password,
	})
	if err != nil {
		log.Println("连接influxDB错误: ", err)
	}
	return cli
}

// ===================================================== 执行SQL语句 =====================================================
func ExecuteSQL(cli client.Client, sql string) (res []client.Result, err error) {
	// Command:         SQL语句（指定下遍参数则下边参数不起作用）
	// Database:        数据库名称
	// RetentionPolicy: 存储策略,
	// Precision:       精度(us,s返回累计数字 空返回日期时间格式)
	selectSQL := client.NewQueryWithRP(sql, "Weather", "1hour", "")
	if response, err := cli.Query(selectSQL); err == nil {
		if response.Error() != nil {
			log.Println("queryDB错误: ", response.Error())
			return res, response.Error()
		}
		res = response.Results
	} else {
		log.Println("queryDB错误: ", err)
		return res, err
	}
	return res, nil
}

// ===================================================== 插入points =====================================================
// 保存数据到数据库需要先获得此结构体各参数，然后调用该结构体的方法 WritesPoints()
type BatchPoints struct {
	// Database is the database to write points to.
	Database string
	// Precision is the write precision of the points, defaults to "ns".
	Precision string
	// RetentionPolicy is the retention policy name of the points, defaults to "autogen".
	RetentionPolicy string
	// Write consistency is the number of servers required to confirm write. ???????????????
	WriteConsistency string
	// Measurement is the database to write points to.
	Measurement string
	// write the point to database
	Point DataPoint
}

type DataPoint struct {
	Tags      map[string]string
	Fields    map[string]interface{}
	Timestamp time.Time
}

//插入数据
func (bpStruct *BatchPoints) WritesPoints(cli client.Client) {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:         bpStruct.Database,         //数据库
		Precision:        bpStruct.Precision,        //精度级别
		RetentionPolicy:  bpStruct.RetentionPolicy,  //存储策略(默认不写为autogen)
		WriteConsistency: bpStruct.WriteConsistency, //
	})
	if err != nil {
		log.Println("NewBatchPoints 错误: ", err)
	}

	pt, err := client.NewPoint(bpStruct.Measurement, bpStruct.Point.Tags, bpStruct.Point.Fields, bpStruct.Point.Timestamp)
	if err != nil {
		log.Println("NewPoint 错误: ", err)
	}

	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Println("Write 错误: ", err)
	}
}
