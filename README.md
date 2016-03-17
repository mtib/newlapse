# NewLapse
This replaces chronolapse on linux

## Usage
```bash
newlapse -capture -rate 10 -folder "recording"
newlapse -crop -folder "recording"
newlapse -convert -fps 20
```
You can also do the following
```bash
newlapse -ccc -rate 5 -fps 10 -folder "tmp"
```
which will capture a picture of the screen every 5 seconds, and save it in the ./tmp folder.
if you CTRL+C it will start cropping the pictures and convert each screen into a mp4 file.

### Features
help text:
```
$ newlapse -help
Usage of newlapse:
  -capture
    	tells newlapse to capture
  -ccc
    	equals '-capture -crop -convert'
  -convert
    	tells newlapse to convert %%ds folders to videos
  -crop
    	tells newlapse to crop
  -folder string
    	which folder to do something with (default "./capture")
  -fps int
    	ffmpeg framerate for videos (default 20)
  -rate int
    	seconds to wait between scrots (default 10)
```
output:
```
capturing into folder: ./capture
3.616E+05 byte/sec = 361.619 kb/sec = 0.362 mb/sec
2.170E+07 byte/min = 21697.140 kb/min = 21.697 mb/min
1GB of storage will be filled in 46.09 Minutes
picture #0000000001 taken
picture #0000000002 taken
picture #0000000003 taken
picture #0000000004 taken
picture #0000000005 taken
picture #0000000006 taken
picture #0000000007 taken
picture #0000000008 taken
^C
cropping folder: ./capture
start cropping (~8 files)
completed cropping
start ffmpeg conversion #1
start ffmpeg conversion #2
start ffmpeg conversion #3
completed conversion
```


## Requirements:
| Programm       | Usage          |
| :------------- | :------------- |
| scrot          | screenshot     |
| imagemagick    | crop           |
| ffmpeg         | convert        |
