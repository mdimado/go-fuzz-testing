#base OSS-Fuzz image for Go projects
FROM gcr.io/oss-fuzz-base/base-builder-go

#project name
ENV PROJECT=fuzztestingingo

WORKDIR /src/$PROJECT

#copy source code into container
COPY . .

#build fuzz target
RUN compile_go_fuzzer fuzztestingingo FuzzEqual fuzz_equal
