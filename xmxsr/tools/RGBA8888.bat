
set TUX_PAKE_PATH="C:\Program Files (x86)\TexturePacker\bin\TexturePacker.exe"



::RGBA8888		ÿ������4���ֽڣ��ʺ�����ƽ̨����win & mac
::PVRTC4		ÿ������2���ֽڣ��ʺ���ios����ƽ̨�Ͳ���Androidƽ̨
::RGBA4444  	ÿ������2���ֽڣ�	
::RGB565		ÿ������2���ֽ�,		
::RGB888		ÿ������2���ֽ�,��Alphaͨ�������ʲ�͸���ı���ͼƬ	
::RGBA5555		ÿ������3���ֽ�,TP��֧�֣�PVR����
::RGBA5551		ÿ������2���ֽ�,


for /f "usebackq tokens=*" %%d in (`dir /s /b *.png`) do (
%TUX_PAKE_PATH% --opt RGBA8888 --no-trim --allow-free-size --disable-rotation --sheet --shape-padding 0 --border-padding 0 "%%~dpnd.pvr.ccz" "%%d"
del "%%d"

)


for /f "usebackq tokens=*" %%a in (`dir /s /b *.pvr.ccz`) do (

set str=%%~na

setlocal enabledelayedexpansion 

set "pre=!str:~0,-4!" 

ren "%%a" "!pre!.png" 

endlocal


)


del out.plist
