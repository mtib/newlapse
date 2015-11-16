#!/usr/bin/python3
import os
import os.path
import sys
import convertToVideo as converter
import takescreenshots as screenshot
import cropscreens
import tkinter
import tkinter.filedialog
import fileio


class App:

    def __init__(self, master):
        frame = tkinter.Frame(master)
        frame.pack()
        self.startrec = tkinter.Button(
            frame, text="Start Recording", command=self.toggleRecorder)
        self.startrec.grid(row=0, sticky="nwse")
        self.crop = tkinter.Button(
            frame, text="Crop Screens", command=self.cropScreens)
        self.crop.grid(row=2, sticky="nwse")
        self.videobtn = tkinter.Button(
            frame, text="Convert to video", command=self.convertDirsToVideo)
        self.videobtn.grid(row=1, sticky="nwse")
        self.recording = False
        self.status = tkinter.Label(frame, text="")
        self.status.grid(row=1, column=1, sticky="nwse")

        self.topbtnfrm = tkinter.Frame(frame)
        self.topbtnfrm.grid(row=0, column=1, sticky="nwse")
        self.settingbtn = tkinter.Button(self.topbtnfrm, text="Settings...")
        self.settingbtn.grid(row=0, column=1, sticky="nwse")
        self.filedialogbtn = tkinter.Button(
            self.topbtnfrm, text="Target Dir...", command=self.calldirdialog)
        self.filedialogbtn.grid(row=0, column=0, sticky="nwse")

    def calldirdialog(self):
        filename = tkinter.filedialog.askdirectory(
            initialdir=fileio.getSetting("target"))
        if filename:
            fileio.saveSetting("target", filename + "/")
        else:
            print("Target not changed")

    def toggleRecorder(self):
        self.recording = not self.recording
        if self.recording:
            print("Now Recording")
            self.startrec.config(text="Stop Recording")
        else:
            print("Stop Recording")
            self.startrec.config(text="Start Recording")

    def cropScreens(self):
        print("Cropping")

    def convertDirsToVideo(self):
        print("Converting to Video")


def main(args):
    root = tkinter.Tk()
    root.title("NewLapse")
    app = App(root)
    root.mainloop()

if __name__ == '__main__':
    from sys import argv
    main(argv)
