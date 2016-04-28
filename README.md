# emd-go

call [Earth mover's distance (c version)](http://ai.stanford.edu/~rubner/emd/default.htm) in Golang

to build, type 'make'

since the change to cgo, add 'GODEBUG=cgocheck=0' to your environment variable before compilation to
avoid a runtime error of 'cgo argument has Go pointer to Go pointer'
