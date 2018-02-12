fork := true

name := "client es"
version := "1.0"
scalaVersion := "2.11.4"

libraryDependencies += "org.apache.spark" %% "spark-core" % "2.1.0"
libraryDependencies += "org.elasticsearch" %% "elasticsearch-spark-20" % "6.1.1"
