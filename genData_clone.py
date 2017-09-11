from random import randint

f = open("/home/user/workspace/open/local/draft/data_clone.csv", "w")
for i in range(100000) :
	f.write("Ben"+str(i)+";Cody"+str(i)+";CC-"+str(i)+";2017-09-10T00:00:00+0200;"+str(randint(1,101))+"\n")