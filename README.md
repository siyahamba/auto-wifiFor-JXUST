# auto-wifiFor-JXUST
# python版本
江西理工大学校园网自动连接脚本，以下教程基于windows系统。需要电脑上存在python3环境，笔者用的是python 3.10 
需要下载对应的依赖包  

在任意位置创建一个文件，输入以下内容： 

$python /'your\ file\ path'/school-wifi.py$  
$pause$  
  
输入完成后保存，将文件命名为jxustCon.bat  
#注：your_file_path 替换为你保留文件school-wifi.py的路径，不能含有中文，存在编码问题  

将.bat文件移动至以下位置：  
     C:\ProgramData\Microsoft\Windows\StartMenu\Programs\Startup  
该位置对计算机所有用户生效  

重启即可验证是否生效  

脚本内容包含连接校园网，校园网验证，检测是否存在网络三部分内容  


#go语言版本
包内main.go为程序源码，在Windows系统上可以直接使用编译好的exe文件  
首次运行文件时需要输入账号密码以及联通方式（电信，移动等），首次输入后即可不再输入  

账号密码存储在与文件同目录下的config.txt(后续更新将替换为config.yml)，如果需要修改密码可以直接在文件中更改或者删除文件重新运行  

使用方式同python版本类似，但是bat文件内的内容更改为：  
$'your\ file\ path'\JXUSTCon-V1.exe$  
$pause$ 
