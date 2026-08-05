package tracker

var values []float64
var windowSize int = 3

func AddValue (value float64) {
values = append(values, value)
if len(values) > windowSize {
values = values[1:]
}}

func GetAverage() float64 {
if len(values) == 0 {
return 0}

total := 0.0
for _, val := range values {
total = total + val
}
return total / float64(len(values))
}