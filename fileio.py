#!/urs/bin/python3
import json
import os

settingfile = "settings.json"


def scriptPath():
    return os.path.dirname(os.path.realpath(__file__)) + "/"


def getSettingFile():
    return scriptPath() + settingfile


def newSettingsFile():
    with open(getSettingFile(), "w") as f:
        f.write(json.dumps({"target": scriptPath()}))


def saveSetting(setting, value):
    obj = {}
    try:
        with open(getSettingFile(), "r") as f:
            obj = json.loads(f.read())
    except:
        newSettingsFile()
    obj[setting] = value
    f = open(getSettingFile(), "w")
    f.write(json.dumps(obj))
    f.close()


def getSetting(setting):
    try:
        with open(getSettingFile(), "r") as f:
            obj = json.loads(f.read())
            return obj[setting]
    except:
        newSettingsFile()
