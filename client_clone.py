import sys

from datetime import datetime
from elasticsearch import Elasticsearch

g_es = Elasticsearch()
g_index = "clones_py"

g_es.indices.delete(index=g_index, ignore=[400, 404])

g_es.indices.create(
    index=g_index,
    body={
        "mappings": {
            "clone" : {
                "properties" : {
                    "firstname" : { "type" : "keyword" },
                    "name" : { "type" : "keyword" },
                    "id" : { "type" : "keyword" },
                    "birthday" : { "type" : "date" },
                    "value" : { "type" : "integer" }
                }
            }
        }
    }
)

def bulkInject(p_es, p_docs):
    l_result = p_es.bulk(p_docs)
    if l_result['errors']:
        print("*** Error: bulk", l_result)
        sys.exit(1)


with open('/home/user/workspace/open/local/draft/data_clone.csv') as c_fp:
    l_num = 0
    l_docs = []
    for c_line in c_fp :
        try : 
            # MCO - SERVICEID - MERCHANT - TAX_INCLUDED_RATING_AMOUNT - TRANSACTION_ID - ERROR_CODE
            l_fields = c_line.split(';')
            l_line_compress = list()

            l_line_compress.append(l_fields[0])         # firstname
            l_line_compress.append(l_fields[1])         # name
            l_line_compress.append(l_fields[2])         # id
            l_line_compress.append(l_fields[3])         # birthday
            l_line_compress.append(l_fields[4][:-1])    # value     remove \n

        except :
            print("PARSE : problem at line " + str(l_num))

#        try :
        l_date = datetime.strptime(l_line_compress[3], "%Y-%m-%dT%H:%M:%S+0200")
        l_doc = {
            'firstname': l_line_compress[0],
            'name': l_line_compress[1],
            'id': l_line_compress[2],
            'birthday': l_date,
            'value': l_line_compress[4]
        }
#        except : 
#            print("INJECTION_ES : problem at line " + str(l_num))
            #raw_input("Press Enter to continue...")
            #print("continue...")
#            continue

        l_docs.append({ "index" : {
            "_index": g_index,
            "_type": "clone"
        } })
        l_docs.append(l_doc)

        l_num += 1

        if 1000 < len(l_docs):
            bulkInject(g_es, l_docs)
            l_docs = []

    if 0 < len(l_docs):
        bulkInject(g_es, l_docs)
        l_docs = []