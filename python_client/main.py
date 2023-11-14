# -*- coding: utf-8 -*-

import requests
import json
import base64
from absl import app,flags

FLAGS = flags.FLAGS

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

flags.DEFINE_string("addr", "http://localhost:8080", "Http service address")
flags.DEFINE_string("mode", "json-img", "Select the startup mode. The optional values are ws, http, form img, and json img")
flags.DEFINE_string("img", "../1.jpg", "Specify image path when sending image messages")
flags.DEFINE_string("wxid", "47331170911@chatroom", "Send message recipient's wxid")

def main(argv):
    if FLAGS.mode == "json-img":
        sendJsonImg()
    elif FLAGS.mode == "form-img":
        sendFormImg()

if __name__ == '__main__':
    app.run(main)