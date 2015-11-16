#!/usr/bin/python3
import os
import sys
import os.path
# ffmpeg -r 1 -pattern_type glob -i 'test_*.jpg' -c:v libx264 out.mp4


def dirtovideo(dirname, framerate=10, glob="*.jpg", videoname=""):
    barecmd = "ffmpeg -y -r {} -pattern_type glob -i '{}' -c:v libx264 {}"
    if len(videoname) > 0:
        cmd = barecmd.format(framerate, dirname + "/" +
                             glob, videoname + ".mp4")
    else:
        cmd = barecmd.format(framerate, dirname + "/" + glob, dirname + ".mp4")
    os.system(cmd)

if __name__ == '__main__':
    dirs = []
    fr = int(input("Frames per Second: "))
    imgtype = ".jpg"
    pwdimgfound = False
    for f in os.listdir():
        if os.path.isdir(f):
            imgfound = False
            for item in os.listdir(f):
                if not imgfound and item.endswith(imgtype):
                    dirs.append(f)
                    imgfound = True
        elif f.endswith(imgtype):
            pwdimgfound = True
    for d in dirs:
        dirtovideo(d, fr)
    if pwdimgfound:
        dirtovideo(".", fr, videoname="pwd_video")
