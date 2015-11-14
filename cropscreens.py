#!/usr/bin/python3
import os
import sys

basedir = "/media/tibyte/Storage/Chronolapse/"

imgFolder = "Work/"
mainFolder = "screen_Main/"
mainWidth = 1920
mainHeight = 1080
tvFolder = "screen_Tv/"
tvWidth = 1280
tvHeight = 768
statFolder = "screen_Stat/"
statWidth = 1280
statHeight = 1024
frmat = "{}x{}+{}+{}"
screens = {
    mainFolder: frmat.format(mainWidth, mainHeight, tvWidth + statWidth, 0),
    tvFolder: frmat.format(tvWidth, tvHeight, statWidth, 0),
    statFolder: frmat.format(statWidth, statHeight, 0, 0)
}


def cut(img):
    for k in screens:
        os.system("convert " + basedir + imgFolder + img +
                  " -crop " + screens[k] + " +repage " + basedir + k + img)


def prepare():
    print("creating folders")
    for d in screens:
        try:
            os.mkdir(basedir + d)
        except:
            pass
    print("created folders")

if __name__ == '__main__':
    prepare()
    imgs = os.listdir(basedir + imgFolder)
    count = 1.0
    length = len(imgs) * 1.0
    for i in imgs:
        cut(i)
        print("Progress: {:5.2%}".format(count / length))
        count += 1.0
