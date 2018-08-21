use 

(cd point && swig -go -c++ -cgo -intgosize 64 point.i)
go build -x ./main.go 
