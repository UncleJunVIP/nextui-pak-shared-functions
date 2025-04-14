package common

import "regexp"

const SDCardRoot = "/mnt/SDCARD"
const RomDirectory = "/mnt/SDCARD/Roms"

var TagRegex = regexp.MustCompile(`\((.*?)\)`)
