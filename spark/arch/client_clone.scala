package pl.japila.spark
import org.apache.spark.SparkContext    
import org.apache.spark.SparkContext._
import org.elasticsearch.spark._        
import org.apache.spark.SparkConf

object Es {
  def main(args: Array[String]) {

	val l_conf = new SparkConf().setAppName("client scala-spark")
    //l_conf.set("es.index.auto.create", "true")
    // l_conf.set("es.write.operation", "index")
    // l_conf.set("spark.default.parallelism", "80")
    // l_conf.set("spark.cores.max", "40")
    // l_conf.set("spark.executor.memory", "22g")
    // l_conf.set("es.scroll.limit", "100")

    //https://www.elastic.co/guide/en/elasticsearch/hadoop/master/configuration.html
    //l_conf.set("es.nodes.client.only", "true")

    l_conf.setMaster("local")
    val sc = new SparkContext(l_conf)

	val inputfile = sc.textFile("/home/user/workspace/open/client_ES/tmp/data_clone.csv")

	inputfile.saveToEs("/clones_spark/clone")

  }
}