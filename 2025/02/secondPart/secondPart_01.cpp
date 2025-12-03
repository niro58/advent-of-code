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


string getInvalidId(string numStr){
    int l = 1;
    while(l < numStr.length()){
      if(numStr.length() % l != 0){
        l += 1;
        continue;
      }
      bool isValid = true;
      string pattern = numStr.substr(0, l);
      for(int j = l; j + l - 1 < numStr.length(); j += l){
        string val = numStr.substr(j, l);
        if(pattern != val){
          isValid = false;
          break;
        }
      }
      if(isValid){
        cout << "IS VALID " << numStr << "\n\n";

        return numStr;
      }
      l += 1;
    }
   return "";
}
int main() {
  string inputContent = getFileContent("input/002.txt");
  string resultContent = getFileContent("result/002.txt");

  auto inputLines = split(inputContent, ",");
  long long int total = 0;
  for(auto v : inputLines ){
    const int delimiterPos = v.find('-');
    const string left = v.substr(0,delimiterPos);
    const string right = v.substr(delimiterPos + 1);
    const long int leftInt =  std::stol(left);
    const long int rightInt = std::stol(right);
 
    for(long int i = leftInt; i <= rightInt;i += 1) {
      const string numStr = std::to_string(i);
      string invalidId = getInvalidId(numStr);
      if(invalidId == ""){
        continue;
      }

      total += std::stol(invalidId);
    }

  }
  const bool equals = resultContent == (std::to_string(total));
  cout << resultContent << "|" << total << "|" << (equals ? "EQUALS" : "RIP");
  return 0;
}