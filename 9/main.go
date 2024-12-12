package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
)

type Sector struct {
	isFile bool
	fileId int
	length int
}

func main() {
	partOne("input")
	partTwo("input")

}

func partOne(path string) {
	disk := parseInputBlock(path)
	disk = defragBlock(disk)
	checksum := calcChecksumBlock(disk)
	fmt.Println(checksum)
}

func partTwo(path string) {
	disk := parseInputSector(path)
	disk = defragSector(disk)
	blockDisk := sectorDiskToBlocDisk(disk)
	checksum := calcChecksumBlock(blockDisk)
	fmt.Println(checksum)
}

func defragSector(disk []Sector) []Sector {
	maxId := findMaxFileId(disk)

	for i := maxId; i >= 0; i-- {
		last, lastI := findSectorWithFileId(disk, i)
		for i := 0; i < lastI; i++ {
			sector := disk[i]
			if !sector.isFile && sector.length >= last.length {
				disk = slices.Delete(disk, lastI, lastI+1)
				disk = slices.Insert(disk, lastI, Sector{false, 0, last.length})
				disk = slices.Insert(disk, i, last)
				disk[i+1].length -= last.length
				if disk[i+1].length == 0 {
					disk = slices.Delete(disk, i+1, i+1)
				}
				break
			}
		}
		//printDiskSector(disk)
	}

	return disk
}

func defragBlock(disk []string) []string {
	for {
		last, lastI := findLastFileBlock(disk)

		end := true
		for i := 0; i < lastI; i++ {
			if disk[i] == "." {
				disk[i] = last
				disk[lastI] = "."
				end = false
				break
			}
		}
		if end {
			break
		}
	}
	return disk
}

func calcChecksumBlock(disk []string) int {
	checksum := 0
	for i, block := range disk {
		if block != "." {
			checksum += utils.StrToInt(block) * i
		}
	}
	return checksum
}

func sectorDiskToBlocDisk(disk []Sector) []string {
	blockDisk := []string{}
	for _, sector := range disk {
		var char string
		if sector.isFile {
			char = utils.IntToStr(sector.fileId)
		} else {
			char = "."
		}
		for i := 0; i < sector.length; i++ {
			blockDisk = append(blockDisk, char)
		}
	}
	return blockDisk
}

func parseInputBlock(path string) []string {
	data := utils.ReadFile(path)
	disk := []string{}
	for i := 0; i < len(data); i += 2 {
		fileId := i / 2
		fileLength := utils.StrToInt(string(data[i]))

		for i := 0; i < fileLength; i++ {
			disk = append(disk, utils.IntToStr(fileId))
		}

		if i+1 < len(data)-1 {
			emptyLength := utils.StrToInt(string(data[i+1]))
			if emptyLength > 0 {
				for i := 0; i < emptyLength; i++ {
					disk = append(disk, ".")
				}
			}

		}
	}
	return disk
}

func parseInputSector(path string) []Sector {
	data := utils.ReadFile(path)
	disk := []Sector{}
	for i := 0; i < len(data); i += 2 {
		fileId := i / 2
		fileLength := utils.StrToInt(string(data[i]))
		disk = append(disk, Sector{true, fileId, fileLength})

		if i+1 < len(data)-1 {
			emptyLength := utils.StrToInt(string(data[i+1]))
			if emptyLength > 0 {
				disk = append(disk, Sector{false, 0, emptyLength})
			}

		}
	}
	return disk
}

func findMaxFileId(disk []Sector) int {
	max := 0
	for _, sector := range disk {
		if sector.isFile && sector.fileId > max {
			max = sector.fileId
		}
	}
	return max
}

func findLastFileBlock(disk []string) (string, int) {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != "." {
			return disk[i], i
		}
	}
	return "FAIL", 0
}

func findSectorWithFileId(disk []Sector, fileId int) (Sector, int) {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i].isFile && disk[i].fileId == fileId {
			return disk[i], i
		}
	}
	return Sector{}, 0
}

func printDiskSector(disk []Sector) {
	for _, sector := range disk {
		var char string
		if sector.isFile {
			char = utils.IntToStr(sector.fileId)
		} else {
			char = "."
		}
		for i := 0; i < sector.length; i++ {
			fmt.Print(char)
		}
	}
	fmt.Println()
}
