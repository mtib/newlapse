#!/usr/bin/python3
import os
import sys
# ffmpeg -r 1 -pattern_type glob -i 'test_*.jpg' -c:v libx264 out.mp4
basedir = "/media/tibyte/Storage/Chronolapse/"
dirs = ["Work/", "screen_Main/", "screen_Stat/", "screen_Tv/"]
fr = int(input("Frames per Second: "))
frmt = "ffmpeg -y -r " + str(fr) + " -pattern_type glob -i '" + \
    basedir + "{}*.png' -c:v libx264 " + basedir + "{}.mp4"
for d in dirs:
    cmd = frmt.format(d, "video_" + d[:-1].lower())
    os.system(cmd)
