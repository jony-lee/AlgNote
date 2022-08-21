import requests
import json
import os
import html2text

level = {
    3:"困难",
    2:"一般",
    1:"简单",
}

# 根据http返回的json内容构造文件名
def get_file_name(x):
    qid = x["stat"]["frontend_question_id"].replace(" ","-")
    title = x["stat"]["question__title"].replace(" ","-").replace("/","")
    return f"{qid:0>4}-{title}"


# 根据http返回的json内容获取问题描述细节
def get_question_detail(title):
    url = "https://leetcode.cn/graphql/"
    headers = {
        'Referer': 'https://leetcode.cn',
        'x-csrftoken': 'Nrux0qFNAA79pprqkj3b7AOcI8eU9kGAIx3dzVE6W84M2Oc60eh5aRXHggXXaDBS',
        'content-type': 'application/json'
    }
    data = {
        "operationName": "questionData",
        "variables": {
            "titleSlug":title,
#             "titleSlug": "check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence"
        },
        "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    categoryTitle\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    exampleTestcases\n    jsonExampleTestcases\n    __typename\n  }\n}\n"
    }
    res = requests.post(url,headers=headers,data = json.dumps(data))
    js = json.loads(res.text)["data"]["question"]
    code = [x["code"] for x in js["codeSnippets"] if x["lang"]=="Go"][0]
    desc = html2text.html2text(js["translatedContent"]) if "<p>" in js["translatedContent"] else js["translatedContent"]
    return (code,desc)

# 根据http返回内容获取go文件里面应该填写的内容
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

# 获取所有题目列表
url = "https://leetcode.cn/api/problems/all/"
res = requests.get(url)
all_problems = json.loads(res.text)
all_problems["stat_status_pairs"]


print(get_file_name(all_problems["stat_status_pairs"][2300]))
print(gen_file_content(all_problems["stat_status_pairs"][10]))

if __name__ == "__main__":
    if not os.path.exists("leetcode"):
        os.mkdir("leetcode")
    read_me_titles = []
    solved = 0
    for item in all_problems["stat_status_pairs"][:10]:
        name = get_file_name(item)
        path = f"leetcode/{name}/"
        # 如果文件存在就根据main文件中的第15行和16行判断这个题目是不是已经完成了,如果完成了就在readme里面设置为done
        status = ""
        if os.path.exists(path):
            go_file = path + "main.go"
            with open(go_file, 'r') as f:
                data = f.readlines()
            if data[5][:4]=="解题思路" and data[6]+data[7] != "\n\n":
                status = "Done"
                solved +=1
        else:
            os.mkdir(path)
            code,desc = get_question_detail(item["stat"]["question__title_slug"])
            with open(path + "desc.md","w") as f:
                f.write(desc)
            with open(path+"main.go","w") as f:
                f.write(gen_file_content(item)+code)
        read_me_titles.append(f"|[{name}]({path}desc.md) |[GO]({path}main.go)|{status}|")

    # 刷新markdown
    read_me_titles.sort(key=lambda x:x)
    total = len(read_me_titles)
    with open("readme.md","w") as f:
        f.write("# 题目列表\n\n")
        f.write(f"![](https://img.shields.io/badge/solved-{solved*100/total}%25(({solved}%2F{total}))-brightgreen)\n")
        f.write("|题目描述|代码|状态|\n|--|--|--|\n")
        f.write("\n".join(read_me_titles))