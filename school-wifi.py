#coding=utf-8
import time
import urllib

import requests
import subprocess

import subprocess


def connect_to_wifi(ssid):
    # 创建一个包含netsh命令的字符串
    command = f'netsh wlan connect name="{ssid}"'

    # 尝试运行命令
    result = subprocess.run(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    print(result)
    # 检查命令是否成功执行
    if result.returncode == 0:
        print("Successfully connected to WiFi.")
        return True
    else:
        print("Failed to connect to WiFi.")
        print("Stderr:", result.stderr)
        return False

    # 使用你的WiFi SSID和密码替换下面的值

# 尝试连接WiFi
ssid = 'JXUST-WLAN'

def ping(host):
    """
    Send a ping request to the given host and return True if the host responds.
    """
    # 构造 ping 命令的字符串
    command = ['ping', host]  # Windows 系统
    # 或者对于 Unix/Linux 系统
    # command = ['ping', '-c', '1', '-W', '1', host]

    # 使用 subprocess 执行命令并获取输出
    result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, universal_newlines=True)
    print(result)
    # 检查输出以确定主机是否响应
    if result.returncode == 1:
        return False
    else:
        return True


def con():
    url = 'http://eportal.jxust.edu.cn:801/eportal/portal/login'
    params = {
        'callback': 'dr1003',
        'login_method': '1',
        'user_account': '一卡通号@联通方式', #选择联通方式，电信：telecom；移动：cmcc；联通：unicom，将对应的单词填入‘联通方式’
        'user_password': '密码',
    }
    headers = {
        'Host': 'eportal.jxust.edu.cn:801',
        'Referer': 'http://eportal.jxust.edu.cn/',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:127.0) Gecko/20100101 Firefox/127.0',
    }

    params2 = {

    }

    response = requests.get(url, params=params, headers=headers)
    print("wifi gets up")
    return(response.text)

host = 'www.baidu.com'
if connect_to_wifi(ssid):
    time.sleep(5)
    con()
    if ping(host):
        print(f"wifi is up.")
    else:
        con()
