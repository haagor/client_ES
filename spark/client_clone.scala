package pl.japila.spark

import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._
import org.elasticsearch.spark._
import org.apache.spark.SparkConf


object Es {
  def main(args: Array[String]) {

    val l_conf = new SparkConf().setAppName("client")
    //l_conf.set("es.nodes", "10.234.172.36:5200")
    //l_conf.set("es.nodes", "192.168.1.31:5200, 192.168.1.31:5201")
    //l_conf.set("es.net.proxy.http.use.system.props", "no")
    l_conf.set("es.write.operation", "index")
    l_conf.set("es.scroll.limit", "100")

    l_conf.setMaster("local")
    val sc = new SparkContext(l_conf)

    // val numbers = Map("one" -> 1, "two" -> 2, "three" -> 3)
    // val airports = Map("arrival" -> "Otopeni", "SFO" -> "San Fran")

    // sc.makeRDD(Seq(numbers, airports)).saveToEs("clones_spark/clone")


    // Read the CSV file
    val csv = sc.textFile("/home/user/workspace/open/client_ES/tmp/data_clone.csv")
    // split / clean data
    val headerAndRows = csv.map(line => line.split(";").map(_.trim))
    // get header
    val header = headerAndRows.first
    // filter out header (eh. just check if the first val matches the first header name)
    val data = headerAndRows.filter(_(0) != header(0))
    // splits to map (header/value pairs)
    val maps = data.map(splits => header.zip(splits).toMap)
    // filter out the user "me"
    //val result = maps.filter(map => map("user") != "me")
    // print result
    maps.saveToEs("clones_spark/clone")

  }
}
