# Lighthouse Scale Fix

A simple utility program to set every SteamVR Ligthouse-tracked device's tracking scale to 1.

This aims to help with lighthouse-tracked controllers and/or trackers being always slightly off but not in the same direction.

Download [here](https://github.com/Louka3000/Lighthouse-Scale-Fix/releases/latest/download/lh_scale_fix.exe).

## Theory

The tracking scale value seems to scale translation from a certain point, which I assume to be the first lighthouse SteamVR initializes.

A scale of 2 would mean that for every 1 meter travelled IRL, 2 meters are travelled in VR.

Values being different for different devices thus mean that they won't align properly.

The solution is thus to set tracking scales to 1. There does not seem to be any downside to doing this; my tracking has been much better ever since doing this.

## Doing it automatically

You can download an executable to do it automatically [here](https://github.com/Louka3000/Lighthouse-Scale-Fix/releases/latest/download/lh_scale_fix.exe).

This edits `lighthouse_scale.json` which it assumes to be at `C:\Program Files (x86)\Steam\steamapps\common\SteamVR\drivers\lighthouse\resources`.

If anything goes wrong, it will have made a .bkp there that you can rename back to revert changes.

## Doing it manually

1. Open `C:\Program Files (x86)\Steam\steamapps\common\SteamVR\drivers\lighthouse\resources\lighthouse_scale.json`.
2. Set the "scale" value to 1 under each device you use.
3. Save the file and restart SteamVR for the change to take effect.

## Building

This assumes you have Golang installed.

To run:

`go run lh_scale_fix.go`

To build into a .exe:

`go build lh_scale_fix.go`
