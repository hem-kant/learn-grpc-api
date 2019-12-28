grpc

vim ~/.bash_profile
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
source ~/.bash_profile

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.

protoc PrimeNumberDecomposition/PrimeNumberDecompositionpb/PrimeNumberDecomposition.proto --go_out=plugins=grpc:.


protoc ComputeAverage/ComputeAveragepb/ComputeAverage.proto --go_out=plugins=grpc:.

protoc Bidi/Bidipb/Bidi.proto --go_out=plugins=grpc:.

protoc Bidi/FindMaximum/findmaximumpb/FindMaximum.proto --go_out=plugins=grpc:.

protoc AdvanceTopics/HandlingErrorAndCodes/calculatorpb/calculator.proto --go_out=plugins=grpc:.

protoc AdvanceTopics/Deadlines/greetpb/greet.proto --go_out=plugins=grpc:.