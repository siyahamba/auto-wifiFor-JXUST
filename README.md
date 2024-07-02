# auto-wifiFor-JXUST
江西理工大学校园网自动连接脚本，以下教程基于windows系统。需要电脑上存在python3环境，笔者用的是python 3.10

在任意位置创建一个文件，输入以下内容：
$python /'your_file_path'/school-wifi.py$ 
$pause\\$
输入完成后保存，将文件命名为jxustCon.bat
#注：your_file_path 替换为你保留文件school-wifi.py的路径，不能含有中文，存在编码问题

将.bat文件移动至以下位置：
C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup
该位置对计算机所有用户生效

重启即可验证是否生效

脚本内容包含连接校园网，校园网验证，检测是否存在网络三部分内容
