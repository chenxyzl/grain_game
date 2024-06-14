#当前目录
CURRENT_DIR := $(shell pwd)

#######################参数#######################
#获取分支名 默认dev分支
ifeq ($(B),)
    A := $(info 缺少分支参数B. 不更新代码, 如果更新代码后编译请使用"make B=分支名字")
else
    A := $(info 本次编译分支为 $(B), 即将更新代码)
    A := $(shell cd .. && git checkout $(B) && git pull --rebase -f && cd $(CURRENT_DIR))
    A := $(info 代码更新完成)
endif

#######################版本#######################
#代码版本
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME := $(shell date "+%F %T")
COMMIT_SHA1 := $(shell git rev-parse HEAD)
COMMIT_TIME := $(shell git log -1 --format="%ci")
COMMIT_LOG := $(shell git log -1 --format="%s")

#编译参数
BUILD_LDFLAG="-X 'grain_game/version.BuildTime=${BUILD_TIME}' -X 'grain_game/version.CommitID=${COMMIT_SHA1}' -X 'grain_game/version.CommitTime=${COMMIT_TIME}' -X 'grain_game/version.Branch=${BRANCH}' "
BUILD_FLAGS=-ldflags $(BUILD_LDFLAG)
BUILD_CMD=go build $(BUILD_FLAGS)

#proto可执行文件位置
export PATH := $(CURRENT_DIR)/tools/:$(PATH)

all: gate home

version:
	@echo CURRENT_DIR:$(CURRENT_DIR)
	@echo BRANCH:$(BRANCH)
	@echo COMMIT_SHA1:$(COMMIT_SHA1)
	@echo COMMIT_TIME:$(COMMIT_TIME)
	@echo COMMIT_LOG:$(COMMIT_LOG)
	@echo BUILD_TIME:$(BUILD_TIME)

gate: version
	$(BUILD_CMD) -o ./bin/gate apps/gate/main.go

home: version
	$(BUILD_CMD) -o ./bin/home apps/home/main.go

clean:
	cd bin && ./run.sh clean

auto:
ifdef T
	@echo $(shell echo "本次为单进程更新:$(T)")
	cd bin && ./run.sh stop $(T) && cd .. && make $(T) && cd bin && ./run.sh restart $(T)
else
	@echo $(shell echo "本次为全进程更新")
	cd bin && ./run.sh stop && ./run.sh clean && cd .. && make all && cd bin && ./run.sh restart
endif

proto: version
	find . -type d -exec sh -c 'cd "{}" && rm -f *.pb.go' \; #查找当前目录和子目录,并删除所有*.pb.go
	protoc -I=./proto/file --go-new_out=paths=source_relative:./proto/gen ./proto/file/inner/*.proto
	protoc -I=./proto/file --go-new_out=paths=source_relative:./proto/gen ./proto/file/ret/*.proto

.PHONY: proto