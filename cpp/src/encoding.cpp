#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include "json/json.h"

using namespace std;

// @autodoc title 向文件写入内容
void writeFileFromString(const string &filename, const string &body)
{
    ofstream ofile(filename);
    ofile<<body;
    ofile.close();
}

// @autodoc title 读取json文件
Json::Value readJsonFile(const string &filename)
{
    ifstream ifile;
    ifile.open(filename);
    Json::CharReaderBuilder ReaderBuilder;
    ReaderBuilder["emitUTF8"] = true;
    Json::Value root;

    string strerr;
    bool ok = Json::parseFromStream(ReaderBuilder, ifile, &root, &strerr);
    if (!ok) {
        cerr << "json解析失败";
    }
    return root;
}