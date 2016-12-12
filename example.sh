# newlapse -capture -rate 3
# newlapse -crop
# newlapse -convert
# ffmpeg -i video_1.mp4 -i video_2.mp4 -filter_complex "[1]scale=iw/2:ih/2 [pip];[0][pip] overlay=main_w-overlay_w:main_h-overlay_h" result.mp4
