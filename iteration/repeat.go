package iteration

const repeatCount = 5
func Repeat(character string) string {
  var repeated string

  for i := 1; i < repeatCount; i++ {
    repeated += character
  }

  return repeated
}
