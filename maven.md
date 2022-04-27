## packaging

### pom
父级的pom文件只作为项目之模块的整合，在`maven install`时不会生成jar/war压缩包。

pom的特点:
1. 可以通过`modules`标签来整合只模块的编译顺序。因此尽量将底层的service放在更先的位置优先加载较为合适。
2. 子模块中共用的依赖项或者将器版本统一写到父级配置中，以便统一管理。
3. 子模块的groupId/artifactId, version 直接从父集继承，减少子模块pom的配置项

> Maven引入依赖使用最短路原则，例如a<–b<–c1.0 ，d<–e<–f<–c1.1，由于路径最短，最终引入的为c1.0；但路径长度相同时，则会引入先申明的依赖

### jar
jar最为常见的打包方式，但pom文件中没有设置`packaging`参数是，默认使用jar方式打包。
这种方式意味着`maven build`时会将这个项目中的所有java文件编译成.class文件，并按照原来的java文件层级结构放置，最终压缩为一个jar文件。

### war
war包和jar包非常相似，同样是编译后的.class文件按层级结构形成文件树后打包成压缩包。不通的是，他会将项目中依赖的所有jar包都放置在WEB-INF/lib这个文件夹下。WEB-INF/classes文件夹仍然放置我们自己带吗编译后的内容。

## relativePath
设定一个空值默认值为../pom.xml 表示将始终从父级仓库中获取，不从本地路径获取

> Maven构建包时的查找顺序为： relativePath元素中的地址/本地仓库/远端仓库。

## dependency的scope

- compile: 编译时有效，打包的成果物也会包含此依赖库。
- provided: 编译和测试是有效，但打包成果物中不回包含此库。
- system: 与`provided`相同，但依赖项不会从maven仓库中下载，而是从本地文件系统中获取，需要配合`systemPath`属性使用。
- runtime: 运行时才会依赖，在编译时不会依赖。
- test: 测试范围有效，在编译和大包的时候不会使用这个依赖项。
- import: 仅使用在`dependencyManagement`中的类型依赖，他表示自定的POM`dependencyManagement`部分中有效的依赖关系列表替换依赖关系。

## import
maven多模块项目结构中，可以使用parent定义父项目，实现从父项目中继承依赖。但maven只能单继承，即一个项目只能使用一个parent标签定义父项目。maven 2.9之后的版本引入了一个新的功能，可以实现依赖上的多重继承。这个功能可以将依赖配置复杂的pom文件拆分成多个独立的pom文件。这样处理可以使得maven的pom配置更加简洁，同时可以复用这些pom依赖。

## dependencyManagement
在项目中申明依赖库的版本号，但只模块依赖这个库时，不需要申明版本号，直接使用父级项目中的版本号。可以方便的统一项目中不同的子模块中依赖库的版本号。


## ref
 [maven scope属性说明](https://www.cnblogs.com/kingsonfu/p/10342892.html)


