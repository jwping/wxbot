#!/bin/bash

rm -f /tmp/.X0-lock

/usr/bin/Xvfb :0 -screen 0 1280x960x24 -nolisten unix -ac +extension GLX +extension RENDER &

if [ -z "${WINEPREFIX}" ]; then
    WINEPREFIX=~/.wine
fi

if [ ! -d "${WINEPREFIX}" ];then
    /usr/bin/winecfg
    sleep 1m
fi

DISPLAY=:0 WINEPREFIX=${WINEPREFIX} WINEDEBUG=-all /usr/bin/wine /home/wxbot/wxbot-sidecar.exe -w 'Z:\home\wxbot\WeChat\WeChat.exe' -b ${WXBOT_ARGS}