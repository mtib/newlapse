#!/usr/bin/python3
import os
import sys

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


def cut(img, screens, folder):
    for k in screens:
        os.system("convert " + imgFolder + img +
                  " -crop " + screens[k] + " +repage " + k + img)


def prepare(screens):
    print("creating folders")
    for d in screens:
        try:
            os.mkdir(d)
        except:
            pass
    print("created folders")

if __name__ == '__main__':
    screens = {
        mainFolder: frmat.format(mainWidth, mainHeight, tvWidth + statWidth, 0),
        tvFolder: frmat.format(tvWidth, tvHeight, statWidth, 0),
        statFolder: frmat.format(statWidth, statHeight, 0, 0)
    }
    imgFolder = sys.argv[1]
    prepare(screens)
    imgs = os.listdir(imgFolder)
    count = 1.0
    length = len(imgs) * 1.0
    for i in imgs:
        cut(i, screens, imgFolder)
        print("Progress: {:5.2%}".format(count / length))
        count += 1.0
