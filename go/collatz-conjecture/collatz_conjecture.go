package collatzconjecture

import "errors"

func CollatzConjecture(ip int) (steps int, err error) {
  if ip < 1 {
    return -1, errors.New("invalid input")
  }

  if ip == 0 {
    return 0, nil
  }

  steps = 0

  for ip != 1 {
    if ip%2 == 0 {
      ip /= 2
    } else {
      ip = 3*ip + 1
    }
    steps++
  }

  return steps, nil
}
