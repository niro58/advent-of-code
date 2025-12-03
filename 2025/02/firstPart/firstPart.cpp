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
      if(numStr.length()% 2 == 1){
        continue;
      }
      const string numStrLeft = numStr.substr(0, numStr.length() / 2);
      const string numStrRight = numStr.substr(numStr.length() / 2 , numStr.length());
      if(numStrLeft == numStrRight){
        total += i;
      }
    }

  }
  const bool equals = resultContent == (std::to_string(total));
  cout << resultContent << "|" << total << "|" << (equals ? "EQUALS" : "RIP");
  return 0;
}