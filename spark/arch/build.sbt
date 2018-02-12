fork := true

name := "client es"
version := "1.0"
scalaVersion := "2.10.4"
libraryDependencies += "org.apache.spark" %% "spark-core" % "1.6.2"
libraryDependencies += "org.elasticsearch" %% "elasticsearch-spark" % "2.4.5"