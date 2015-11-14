#!/usr/bin/python3
import os
import sys
import time
imgDir = "/media/tibyte/Storage/Chronolapse/Work/"

try:
    os.mkdir(imgDir[:-1])
except:
    pass


def main():
    interval = float(input("Screenshot interval (in seconds): "))
    num = 1
    while True:
        imgname = str(int(time.time() * 1000)) + ".png"
        os.system("scrot -q 85 " + imgDir + imgname)
        print("Screenshot: " + imgname, "num:", num)
        time.sleep(interval)
        num += 1

if __name__ == '__main__':
    main()
