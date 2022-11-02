封装了spr机器人的基础抽象类
有2个作用：
（1）spr_robot_go 机器人依赖的公共库，robot机器人需要继承该仓库的abstract.AbstractRobot，scene需要继承该仓库的abstract.AbstractRobotMgr
（2）打包为spr.so也需要依赖改库，打包为plugin需要对外开发公共api，一些api需要的参数或者返回值都是该仓库定义的类型

仓库地址：https://github.com/tester-rep/spr-go.git

go mod方式依赖本仓库的方式：在go.mod中require(github.com/tester-rep/spr-go v0.1.3)

更新本仓库的方式：
（1）修改本地代码
（2）提交：
git add .
git commit -m"new feature"
git push
(3)新增分支：
git tag --list
git tag v0.1.4
git push --tags
