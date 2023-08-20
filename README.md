# Ligthouse Scale Fix

A simple utility program to set every SteamVR Ligthouse-tracked device's tracking scale to 1.

### Theory

The tracking scale value seems to scale translation from a certain point, which I assume to be the first ligthouse SteamVR initializes.
A scale of 2 would mean that for every 1 meter traveled IRL, 2 meters are traveled in VR.
For some reason, Valve decided to set random device scales which can cause issues with tracking in big playspaces.

### Before running

This program edits `lighthouse_scale.json` which it assumes to be at `C:\Program Files (x86)\Steam\steamapps\common\SteamVR\drivers\lighthouse\resources`.
If anything goes wrong, it will have made a .bkp there that you can rename back.

### Running and building

If you're not a developer, just download a release and launch the .exe!

We assume you have Go installed.

To run:

`go run lh_scale_fix.go`

To build into a .exe:

`go build lh_scale_fix.go`