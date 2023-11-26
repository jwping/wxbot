# -*- coding: utf-8 -*-
import os
import requests
import json
import base64
from absl import app,flags

FLAGS = flags.FLAGS


def sendtxtmsg():
    data = {
        "wxid": FLAGS.wxid,
        "content": "测试内容\nhello world!",
    }
    ret = requests.post(url=FLAGS.addr + "/sendtxtmsg", json=data)
    if ret.status_code == 200:
        print("success: %s" % (ret.text))
    else:
        print("faild: %s" % (ret.text))

def sendFormImg():
    data = {"wxid":FLAGS.wxid}
    files = {
                'image': open(FLAGS.img, 'rb')
            }
    ret = requests.post(url=FLAGS.addr + "/sendimgmsg", data=data, files=files)
    if ret.status_code == 200:
        print("success: %s" % (ret.text))
    else:
        print("faild: %s" % (ret.text))


def sendJsonImg():
    with open(FLAGS.img, 'rb') as file:
        data = {
            'wxid': FLAGS.wxid,
            'image': base64.b64encode(file.read()).decode(),
        }

        ret = requests.post(url=FLAGS.addr + "/sendimgmsg", json=data)
        if ret.status_code == 200:
            print("success: %s" % (ret.text))
        else:
            print("faild: %s" % (ret.text))

def sendFormFile():
    data = {"wxid":FLAGS.wxid}
    files = {
                'file': open(FLAGS.file, 'rb'),
                'clear': "false",
            }
    ret = requests.post(url=FLAGS.addr + "/sendfilemsg", data=data, files=files)
    if ret.status_code == 200:
        print("success: %s" % (ret.text))
    else:
        print("faild: %s" % (ret.text))


def sendJsonFile():
    with open(FLAGS.file, 'rb') as file:
        data = {
            'wxid': FLAGS.wxid,
            'file': base64.b64encode(file.read()).decode(),
            'filename': os.path.basename(FLAGS.file)
        }

        ret = requests.post(url=FLAGS.addr + "/sendfilemsg", json=data)
        if ret.status_code == 200:
            print("success: %s" % (ret.text))
        else:
            print("faild: %s" % (ret.text))

flags.DEFINE_string("addr", "http://localhost:8080", "Http service address")
flags.DEFINE_string("mode", "json-file", "Select the startup mode. The optional values are ws, http, form-img, json-img, form-file and json-file")
flags.DEFINE_string("img", "../public/1.jpg", "Specify image path when sending image messages")
flags.DEFINE_string("file", "../public/1.txt", "Send file message specifying file path")
flags.DEFINE_string("wxid", "47331170911@chatroom", "Send message recipient's wxid")

def main(argv):
    if FLAGS.mode == "json-img":
        sendJsonImg()
    elif FLAGS.mode == "form-img":
        sendFormImg()
    elif FLAGS.mode == "json-file":
        sendJsonFile()
    elif FLAGS.mode == "form-file":
        sendFormFile()

if __name__ == '__main__':
    app.run(main)