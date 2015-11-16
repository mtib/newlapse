#!/usr/bin/python3
import os
import sys
import time
import platform
osname = platform.uname().system


def takeScreenshot(filename):
    if osname == "Darwin":
        os.system("screencapture -x -T 0 -t jpg " + filename)
        os.system("convert " + filename + " -scale 50% tmp_" + filename)
        os.remove(filename)
        os.rename("tmp_" + filename, filename)
    elif osname == "Linux":
        os.system("scrot -q 85 " + filename)
    else:
        print("Taking screenshots doesn't work under windows")
        sys.exit()
    print("New screenshot:", filename)


def main():
    print("Working in", os.path.abspath(os.curdir))
    interval = float(input("Screenshot interval (in seconds): "))
    prefix = input("Screenshot prefix: ")
    while True:
        imgname = prefix + str(int(time.time() * 1000)) + ".jpg"
        takeScreenshot(imgname)
        time.sleep(interval)


if __name__ == '__main__':
    main()
