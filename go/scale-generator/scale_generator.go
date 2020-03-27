package scale

import (
  "fmt"
  "strings"
  "unicode"
)

func getPitchFromTonic(tonic string, isFlat bool) []string {
  pitches := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

  pitchFlat := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
 
  fmt.Println(isFlat)
  if isFlat {
    pitches = pitchFlat
  }
  index := -1

  for i, v := range pitches {
    if v == strings.ToUpper(tonic) {
      index = i
    }
  }

  newPitch := make([]string, 0)
  newPitch = append(newPitch, pitches[index:]...)
  newPitch = append(newPitch, pitches[:index]...)

  return newPitch
}

func Scale(tonic, intervale string) []string {
  pitches := getPitchFromTonic(tonic, unicode.IsLower(rune(tonic[0])))
  
  if intervale == "" {
    return pitches
  }

  intervales := strings.Split(intervale, "")
  fmt.Println(intervales)
 
  retPitch := []string{pitches[0]}

  index := 0
  for _, v := range intervales {
    if v == "M" {
      index += 2
    } else if v == "m" {
      index += 1
    } else if v == "A" {
      index += 3
    }
    if index > len(pitches) {
      index -= len(pitches)
    }
    retPitch = append(retPitch, pitches[index])
  }

  return retPitch
}
