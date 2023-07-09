#include <iostream>
#include <fstream>

using namespace std;

// @autodoc title 创建文件
int createFile()
{
    ofstream oFile;
    //不存在则新建文件
    oFile.open("test1.txt", ios::app);
    if (!oFile)  //true则说明文件打开出错
        cout << "error 1" << endl;
    else
        oFile.close();
    return 0;
}
