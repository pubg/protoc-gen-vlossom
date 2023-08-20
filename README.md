# protoc-gen-vlossom

Vlossom의 Dynamic Form을 만드는 proto plugin

## Getting Started
1. protoc-gen-vlossom을 설치합니다.
```shell
go install github.com/pubg/protoc-gen-vlossom/cmd/protoc-gen-vlossom@main
```
2. proto/options.proto를 Workdir에 복사합니다. 
3. 아래 내용대로 example.proto 파일을 생성합니다. 
```
syntax = "proto3";
package schema;
import "options.proto";
option (pubg.vlossom.file) = {expose: true, entrypointMessage: "Values"};

message Values {
  string sample_input = 1 [(pubg.vlossom.field) = {component: Input, input: {type: text, max: 10}}];
  MyEnum my_enum = 2;
}

enum MyEnum {
  FOO = 0;
  BAR = 1;
  BAZ = 2;
}
```
4. 다음 명령어를 실행합니다.
```shell
protoc \
  --vlossom_out=./ \
  -I ./ \
  ./example.proto
```

#### Workdir 상태
options.proto
example.proto


2. example.proto에 아래 내용을 채웁니다.

## Protobuf AST
![ast](./ast.png)

## Message Features
- [ ] oneof: 어떻게?
- [x] map: Json Editor?
- [x] scala: 주 과제
- [ ] repeated: Scala, messsage의 list형, 어떻게?
- [x] optional 
- [x] expose
- [x] message in message

## Scala Field Features
- [ ] select with external source: 툴파트랑 필드 이름 논의해야 함
- [x] string options: min, max, regex?
- [x] number options: ?
- [x] boolean type: checkbox or toggle
- [x] enum type: select or radio set

## Scala Type
1. String
2. Number
3. Enum
4. Boolean

## Well-known type
- [ ] Any
- [ ] timestamp
- [ ] duration
- [ ] k8s.io.apimachinery.pkg.apis.util.v1.IntOrString
- [ ] k8s.io.api.pkg.core.v1.Volume
- [x] map
- [x] boolean-repeated
- https://protobuf.dev/reference/protobuf/google.protobuf/

## 추후 과제
- [ ] 2차원 Location Override
- [x] Property Override
- [ ] protobuf package name change to [pubg.vlossom]

# How run without protoc (Standalone mode)
1. sh examples/debug_input_generate.sh 
2. cat examples/input.dump | go run main.go

protoc 없이 실행하기 위해선 GeneratedRequest(protobuf)를 생성해야 한다.

GeneratedRequest는 바이너리 포맷이라 텍스트로 관리하기 어렵다.

`debug_input_generate.sh` 스크립트를 사용해 example.proto로 부터 GeneratedRequest를 추출한다.

stdin으로 입력은 `cat examples/input.dump |` 또는 Goland에서 `다음 위의 입력 리디렉션`을 활성화하면 된다.

Happy Debugging!

# How run with protoc (Plugin mode)
1. sh examples/generate.sh
