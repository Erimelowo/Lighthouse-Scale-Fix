# Lighthouse Scale Fix

A simple utility program to set every SteamVR Ligthouse-tracked device's tracking scale to 1.

This aims to help with lighthouse-tracked controllers and/or trackers being always slightly off but not in the same direction.

**This may not work for everyone and could lead to worst results (idk why). Only do it if you have the above tracking problem. If you have worse results, just rename the .bkp back that the program created and try modifying the scales manually.**

Download [here](https://github.com/Louka3000/Lighthouse-Scale-Fix/releases/latest/download/lh_scale_fix.exe).

## Theory

The tracking scale value allows to scale translation from a certain point in space. A scale of 2 would mean that for every 1 meter travelled IRL, 2 meters are travelled in VR.

Different values need different scales, but for some setups, the default values may not be good for whatever reason.

The solution is thus to edit the values yourself.

## Doing it automatically

You can download an executable to do it automatically [here](https://github.com/Louka3000/Lighthouse-Scale-Fix/releases/latest/download/lh_scale_fix.exe).

This edits `lighthouse_scale.json` which it assumes to be at `C:\Program Files (x86)\Steam\steamapps\common\SteamVR\drivers\lighthouse\resources`.

If anything goes wrong, it will have made a .bkp there that you can rename back to revert changes.

## Doing it manually

**Make sure to backup**

1. Open `C:\Program Files (x86)\Steam\steamapps\common\SteamVR\drivers\lighthouse\resources\lighthouse_scale.json`.
2. Set the "scale" value to 1 under each device you use (note: you can use different values if need be).
3. Save the file and restart SteamVR for the change to take effect.

## Building

This assumes you have Golang installed.

To run:

`go run lh_scale_fix.go`

To build into a .exe:

`go build lh_scale_fix.go`
