
def average(p_tab):
	if len(p_tab) == 0 :
		return ""
	l_tab = sorted(p_tab)[1:-1]
	return str(sum(l_tab)/len(l_tab))

g_f = open("/Users/simon/workspace/client_ES/tmp/timeAverage", "w")
with open('/Users/simon/workspace/client_ES/tmp/time') as c_fp:
	l_average = []
	for c_line in c_fp :
		if c_line == "\n" or c_line == "" :
			pass
		elif c_line[0] == "#" :
			g_f.write(average(l_average) + "\n")
			g_f.write(c_line[:-1] + "\t:\t")
			l_average = []
		else :
			l_average.append(float(c_line))
	g_f.write(average(l_average))
