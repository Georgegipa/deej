# About this fork

This fork improves upon the work by [Miodec](https://github.com/Miodec/deej) which adds support for remappable buttons by simplifying the process of adding new buttons. It also adds support for an oled display, that displays the changed volume and the time while idling.

# Additions
* Support for oled display that shows the changed volume with the corresponding session name
* Display time while idling
* Better support for remappable buttons 
    * The names of the supported buttons are defined [here](./pkg/deej/kbmap.go)
* Support for button combinations

# How to

 - Upload the Arduino code to your board (be sure to change the pin definitions).
 - Download the `deej-release.exe` file and `config.yaml` from the [release section](https://github.com/Georgegipa/deej/releases/tag/oled) (`deej-dev` will show a debug console window when launched)
 - If you run into any issues, feel free to ask in the [Deej Discord Server](https://discord.gg/nf88NJu)

# Special Thanks

- omriharel for the original [deej](https://github.com/omriharel/deej)
- Miodec for adding [remappable button](https://github.com/Miodec/deej) support
