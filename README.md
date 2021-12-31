pf 全称 protocol formatter，可以格式化 proto 文件，也可以只检查输入文件的是否符合格式化标准。

# 使用方法
1. 格式化后输出到标准输出
> pf hello.proto

2. 检查是否符合格式化规范，diff 输出到标准输出
> pf -l hello.proto
```
message recipient {
-    optional int32 id = 1;
+  optional int32 id = 1;
   optional string name = 2;
   optional string email = 3;
   optional int32 report_type = 4;
```

3. 原地格式化文件
> pf -w hello.proto


