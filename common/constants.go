package common

import "regexp"

const SDCardRoot = "/mnt/SDCARD"
const RomDirectory = "/mnt/SDCARD/Roms"
const CollectionDirectory = "/mnt/SDCARD/Collections"

var TagRegex = regexp.MustCompile(`\((.*?)\)`)
