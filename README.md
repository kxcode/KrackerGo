# KrackerGo

PBKDF2 SHA256 Hash Cracker for Gogs

# Usage

```
brew install golang
export GOPATH=`pwd`
cd src
go run kracker.go
```

```
+------+------------------------------------------------------------------------------------------------------+------------+------------+
| name | passwd                                                                                               | rands      | salt       |
+------+------------------------------------------------------------------------------------------------------+------------+------------+
| root | 30aa3f0f7ffe440cd8c808154fca0f4a555e896de8ba7e285ad0c125eed81d47b89fdbdf14abb83bcb7a5a4aa57dcbd138af | U059wmVWer | LU5pzgLZ6l |
| test | 9862375025386c681e0c8d8cf05e3e536b307dba7c068112650429723f1cd326883af3600f1cc5c7386ae52e426865a75618 | 4O2dGQq6X1 | tegrBTKHrp |
+------+------------------------------------------------------------------------------------------------------+------------+------------+
测试的数据，密码分别是toor 和 test
```

运行效果：
```
$ go run kracker.go
[Cracked]  root toor 30aa3f0f7ffe440cd8c808154fca0f4a555e896de8ba7e285ad0c125eed81d47b89fdbdf14abb83bcb7a5a4aa57dcbd138af LU5pzgLZ6l
pwdfile read ok!
```


Hash数据放在`user.csv`中，分三列，分别是username/hash/salt。

密码放在password.txt中。


# Gogs加密算法

<https://github.com/gogs/gogs/issues/42>
2014年 Gogs的加密算法升级为 PBKDF2 + hmac + sha256


# Hashcat

Hashcat支持破解PBKDF2 SHA256，但是参数格式没弄明白，懒得折腾，所以写了KrackerGo临时用用。

Hashcat命令相关信息如下：

`hashcat -m 10900 --force hash.txt password.txt`

hash.txt内容：

`sha256:10000:TFU1cHpnTFo2bAo=:MKo/D3/+RAzYyAgVT8oPSlVeiW3oun4oWtDBJe7YHUe4n9vfFKu4O8t6WkqlfcvROK8=`

salt和hash都需要转换成base64格式，可以参考这篇：

https://github.com/hashcat/hashcat/issues/1583
