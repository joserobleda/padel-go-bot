package main

type track struct {
	id   string
	name string
}

func getTracks() []track {
	if domain == "canaldeisabel.padelclick.com" {
		tracks := make([]track, 5)

		tracks[2] = track{"1480", "#4"}
		tracks[4] = track{"1477", "#1"}
		tracks[1] = track{"1479", "#3"}
		tracks[0] = track{"1478", "#2"}
		tracks[3] = track{"1481", "#5"}

		return tracks
	}

	return make([]track, 0)
}
