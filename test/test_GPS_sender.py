data=[]
with open(r'log.txt') as file:
    for i in file:
        data.append(i.strip('\n').split(","))

import urllib.request
import time
while(1):
    for i in range(len(data)):
        url = "http://192.168.145.223:8080/setPosition"
        param = {
            "id":1,
            "latitude":float(data[i][0]),
            "longitude":float(data[i][1])
            }
        url += "?{}".format( urllib.parse.urlencode(param) )
        #API実行
        with urllib.request.urlopen(url) as res:
            html = res.read().decode("utf-8")
            print(url)
        
        if i + 1 != len(data):
            before=int(data[i][2].split(":")[2])
            after=int(data[i + 1][2].split(":")[2])
            if after <10:
                after += 60
            time.sleep(after - before)
