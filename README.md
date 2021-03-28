# TubeCheck
一个用于检测/诊断Youtube地域信息的脚本

![image](https://user-images.githubusercontent.com/13616352/112748155-97658a00-8fec-11eb-82a2-1839781ce968.png)

**New Feature:**
1. 支持双栈网络测试
2. 显示当前网络到Youtube的连接模式
3. 在`Google Global Cache`连接模式下，会提供ISP信息
4. 支持获取当前网络所访问Youtube视频节点的地域信息
4. 如果地域信息可用，会显示Youtube识别到的地区信息

此为Beta版本，可能会有一些地域信息数据的缺失，您可以发Issue协助完善它

## 使用方法

支持IPv4网络的机器：

`wget -O tubecheck https://github.com/sjlleo/TubeCheck/releases/download/1.0Beta/tubecheck_1.0beta_linux_amd64 && chmod +x tubecheck && clear && ./tubecheck`

仅支持IPv6网络的机器：

`wget -O tubecheck https://cdn.jsdelivr.net/gh/sjlleo/TubeCheck/CDN/tubecheck_1.0beta_linux_amd64 && chmod +x tubecheck && clear && ./tubecheck`
