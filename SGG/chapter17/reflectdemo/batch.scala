package spark

import org.apache.spark.{SparkConf, SparkContext}
import org.junit.Test

class batch {
  @Test
  def batch1: Unit = {
    val conf = new SparkConf().setMaster("local[8]").setAppName("batch")
    val sc = new SparkContext(conf)

    val source = sc.textFile("hdfs://192.168.176.132/bigwork2/all.csv")

    val map = source.map(item => (item.split(",")(1), item.split(",")(2),item.split(",")(3).split(" ")(0)))
      .map(item => (item,1))

    val reduce = map.reduceByKey((cur , agg) => cur + agg)

    val res = reduce.collect()
    res.foreach(item => println(item))


  }
}
