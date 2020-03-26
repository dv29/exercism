package listops

type IntList []int
type binFunc func (int, int) int
type predFunc func (int) bool
type unaryFunc func (int) int

func (l IntList) Foldr(fn binFunc, initial int) int {
  for i := len(l)-1; i >= 0; i-- {
    initial = fn(l[i], initial)
  }
  return initial
}

func (l IntList) Foldl(fn binFunc, initial int) int {
  for _, v := range l {
    initial = fn(initial, v)
  }
  return initial
}

func (l IntList) Append(listToAppend IntList) IntList {
  newList := make(IntList, len(l)+len(listToAppend))
  lLen := len(l)

  for i, v := range l {
    newList[i] = v
  }
  for i, v := range listToAppend {
    newList[i+lLen] = v
  }
  return newList
}

func (l IntList) Concat(args []IntList) IntList {
  newList := make(IntList, 0)

  newList = append(newList, l...);

  for _, v := range args {
    newList = append(newList, v...)
  }
  return newList
}

func (l IntList) Filter(fn predFunc) IntList {
  newList := make(IntList, 0)

  for _, v := range l {
    if fn(v) {
      newList = append(newList, v)
    }
  }

  return newList
}

func (l IntList) Length() int {
  return len(l)
}

func (l IntList) Map(fn unaryFunc) IntList {
  newList := make(IntList, len(l))

  for i, v := range l {
    newList[i] = fn(v)
  }

  return newList
}

func (l IntList) Reverse() IntList {
  newList := make(IntList, len(l))
  counter := 0

  for i := len(l)-1; i >= 0; i-- {
    newList[counter] = l[i]
    counter++
  }

  return newList
}
