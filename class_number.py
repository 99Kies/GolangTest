import requests
import re
import json
from requests.exceptions import RequestException
import time

def get_one_page(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            response.encoding = response.apparent_encoding
            return response.text
        return None
    except RequestException:
        return None

def parse_one_page(html):
    pattern = re.compile('<tr>.*?<td><a.*?href.*?>(.*?)</a></td>.*?<td>(.*?)</td>.*?<td>(.*?)</td>.*?</tr>',re.S)
    items = re.findall(pattern,html)
    for item in items:
        yield {
            '学号':item[0],
            '姓名':item[1],
            '班级':item[2]
        }

def write_to_file(content):
    with open('./r.txt','a',encoding='utf-8') as f:
        f.write(json.dumps(content,ensure_ascii=False) + '\n')

def main(offset):
    url = 'http://www.xuxianming.cn/look.asp?keywords=&page=' + str(offset)
    html = get_one_page(url)
    for item in parse_one_page(html):
        print(item)
        write_to_file(item)

if __name__ == '__main__':
    for i in range(1,828):
        main(i)
        time.sleep(1)

