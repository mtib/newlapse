# NewLapse
This aims to be a python3 replacement for chronolapse on linux.

## Usage
```
python3 takescreenshots.py  # take fullscreen screenshot with interval
python3 cropscreens.py      # hardcoded to fit my setup
python3 convertToVideo.py   # hardcoded to fit my setup

# my setup
# 3 screens: 1280x1024, 1280x768, 1920x1080
# high priority to be able to change this via a settings file
```

### Todo:
- [ ] relative file path
- [ ] multiplatform recoder
- [ ] replace imagemagick with pil(low)

### Deps:
| Programm       | Usage          | Replacement    |
| :------------- | :------------- | :------------- |
| scrot          | screenshot     | multiplatform options |
| imagemagick    | crop           | pil or pillow  |
