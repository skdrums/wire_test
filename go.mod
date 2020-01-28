module github.com/skdrums/wire_test

go 1.13

replace github.com/tensorflow/tensorflow/tensorflow/go/core => ./tf_grpc/proto/tensorflow/core

require (
	github.com/golang/protobuf v1.3.2
	github.com/tensorflow/tensorflow/tensorflow/go/core v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.26.0
)
