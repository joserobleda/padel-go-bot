package main

import "strings"

type track struct {
	id   string
	name string
}

func getTracks() []track {
	if strings.Contains(domain, "canaldeisabel") {
		tracks := make([]track, 5)

		tracks[0] = track{"1480", "#4"}
		tracks[1] = track{"1477", "#1"}
		tracks[2] = track{"1479", "#3"}
		tracks[3] = track{"1478", "#2"}
		tracks[4] = track{"1481", "#5"}

		return tracks
	}

	if strings.Contains(domain, "ocioydeportecanal") {
		tracks := make([]track, 8)

		tracks[0] = track{"1477", "#1"}
		tracks[1] = track{"1478", "#2"}
		tracks[2] = track{"1479", "#3"}
		tracks[3] = track{"1480", "#4"}
		tracks[4] = track{"1481", "#5"}
		tracks[5] = track{"1482", "#6"}
		tracks[6] = track{"1483", "#7"}
		tracks[7] = track{"1484", "#8"}

		return tracks
	}

	return make([]track, 0)
}
