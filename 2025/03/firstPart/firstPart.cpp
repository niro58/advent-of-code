#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <stdexcept>
using namespace std;

string getFileContent(string path) {
  std::ifstream file(path);
  std::ostringstream content;

  content << file.rdbuf();

  return content.str();
}

std::vector<std::string> split(const std::string& s, const std::string& delimiter) {
    std::vector<std::string> tokens;
    size_t pos = 0;
    size_t start = 0;
    std::string token;
    while ((pos = s.find(delimiter, start)) != std::string::npos) {
        token = s.substr(start, pos - start);
        tokens.push_back(token);
        start = pos + delimiter.length();
    }
    tokens.push_back(s.substr(start));

    return tokens;
}



int main() {
  string inputContent = getFileContent("input/002.txt");
  string resultContent = getFileContent("result/002.txt");

  auto inputLines = split(inputContent, "\n");
  int total = 0;
  for(auto v : inputLines ){
    cout << v << "\n";
    char a = 0;
    char b = 0;
    int i = 0;
    while(i < v.length()){
      const int cNum = std::stoi(to_string(v[i]));
  
      if(i + 1 < v.length() && cNum > a){
        a=cNum;
        b=v[i+1];
      }else if (cNum > b){
        b = cNum;
      }
      i += 1;
    }
    cout << (a - '0') << "|" << (b - '0') << "\n";
    total += (a - '0') * 10 + (b - '0');
  }
  const bool equals = resultContent == (std::to_string(total));
  cout << resultContent << "|" << total << "|" << (equals ? "EQUALS" : "RIP");
  return 0;
}