import requests
import json
url = "https://leetcode.cn/api/problems/all/"
res = requests.get(url)
js = json.loads(res.text)
js["stat_status_pairs"]

level = {
    3:"困难",
    2:"一般",
    1:"简单",
}
def get_file_name(x):
    qid = x["stat"]["frontend_question_id"].replace(" ","-")
    title = x["stat"]["question__title"].replace(" ","-").replace("/","")
    return f"{qid:0>4}-{title}.go"
def gen_file_content(x):
    
    return f"""package leetcode
/*
标题 : {x["stat"]["question__title"]}
链接 : https://leetcode.cn/problems/{x["stat"]["question__title_slug"]}/
难度 : {level[x["difficulty"]["level"]]}
解题思路:


注意事项:

*/
"""
print(get_file_name(js["stat_status_pairs"][2300]))
print(gen_file_content(js["stat_status_pairs"][10]))

import os
os.mkdir("leetcode")

for item in js["stat_status_pairs"]:
    file = f"leetcode/{get_file_name(item)}"
    if os.path.exists():
        continue
    with open(file,"w") as f:
        f.write(gen_file_content(item))