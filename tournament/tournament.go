package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Stats struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}
type QuickStats struct {
	Team string
	P    int
}

func Tally(reader io.Reader, writer io.Writer) error {
	read, _ := ioutil.ReadAll(reader)
	// split matches
	var matches [][]string
	splits := strings.Split(string(read), "\n")
	if len(splits) < 3 {
		return errors.New("invalid match input 2")
	}
	for _, r := range splits {
		if r == "" {
			continue
		}
		m := strings.Split(r, ";")
		if len(m) < 3 {
			continue
		}
		matches = append(matches, m)
	}
	// parse matches -> build stats struct
	statsMap := make(map[string]Stats)
	for _, match := range matches {
		result := match[2]

		if result != "win" && result != "loss" && result != "draw" {
			return errors.New("result is bullshit")
		}

		for i := 0; i < 2; i++ {
			team := match[i]
			// check if map contains team
			if _, ok := statsMap[team]; !ok {
				// add team to map
				statsMap[team] = Stats{}
			}
			// set stat for result
			stats := statsMap[team]
			// calculate results
			statsMap[team] = calculatePoints(stats, result, i)
		}

	}
	var points []QuickStats

	for k, v := range statsMap {
		q := QuickStats{k, v.P}
		points = append(points, q)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].P == points[j].P {
			return points[i].Team < points[j].Team
		}
		return points[i].P > points[j].P
	})

	// build stats table -> return result
	header := `Team                           | MP |  W |  D |  L |  P`
	fmt.Fprintln(writer, header)
	for _, quickStat := range points {
		for k, v := range statsMap {
			if v.P == quickStat.P && k == quickStat.Team {
				fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", k, v.MP, v.W, v.D, v.L, v.P)
			}
		}
	}

	return nil
}

func calculatePoints(stats Stats, result string, teamIndex int) Stats {
	if teamIndex == 1 && result != "draw" {
		if result == "win" {
			result = "loss"
		} else {
			result = "win"
		}
	}
	stats.MP += 1
	switch result {
	case "win":
		stats.W += 1
		stats.P += 3
	case "draw":
		stats.D += 1
		stats.P += 1
	case "loss":
		stats.L += 1
	}
	return stats
}
