# Retro Game Randomizer

## Table of Contents

## Purpose

The Randomizer is a terminal based application used for selecting a random retro video game title either from:

- 2 or More Manually Entered Titles
- A Non-Comprehesive List of Included Titles
- A Custom List in .csv File Format

For the purpose of the included list, a retro game is defined as:

- Any console title prior to 7th Generation consoles (PS3, 360, Wii)
- Any handheld title prior to 7th Generation handhelds (NDS, PSP)
- Any PC title released prior to the year 2000

The included list currently supports all US released titles for the following consoles:

- Atari
  - Atari 2600
  - Atari 5200
  - Atari 7800
  - Atari Lynx
- Nintendo
  - Super Nintendo
  - Nintendo 64
  - Gameboy
  - Gameboy Color
  - Gameboy Advance
- Sega
  - Sega Master System
  - Sega Genesis
  - Sega 32X
  - Game Gear
- Sony
  - Playstation
- Misc
  - Coleco Vision
  - MSX
  - Neo Geo Pocket Color

## Installation

### Pre Compiled Binaries

### Build From Source

For Arch Linux

```bash
sudo pacman -S go
git clone https://github.com/daneburchette/randomizer.git
./scripts/build.sh
```

The build script will build linux and windows binaries inside the randomizer/bin directory.

## Configuration

The first time the Randomizer is run in CSV mode, if a csv file is not found, a default will be generated with the games for the listed systems. However, this file can be replaced with a custom csv file, or altered after generation to add or remove. The expected format is as follows:

```csv
"title","console"
"Game 1 Name","Console 1 Name"
"Game 2 Name", "Console 2 Name"
...
```
