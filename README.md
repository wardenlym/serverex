# ServerEx

安装mongodb后，clone项目到 gitlab.com/megatech/serverex

    make

协议代码生成

    cd apps/excalibur && make codegen

单元测试

    cd apps/excalibur && make test

beta测试环境上线

    make deploy
