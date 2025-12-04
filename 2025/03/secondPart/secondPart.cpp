#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <stdexcept>
#include <cmath>

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
  long long int total = 0;
  const int TOTAL_BATTERIES = 12;
  for(auto v : inputLines ){
    cout << v << "\n";
    int batteries[TOTAL_BATTERIES] = {};
    int i = 0;
    while(i < v.length()){
      const int cNum = std::stoi(to_string(v[i])) - '0';
      for(int j = 0 ; j < TOTAL_BATTERIES; j++){
        if(i + (TOTAL_BATTERIES - 1 - j) < v.length() && cNum > batteries[j]){
          batteries[j] = cNum;
          for(int x = j + 1; x < TOTAL_BATTERIES; x++){
            batteries[x] = 0;
          }
          break;
        }
      }
      i += 1;
    }
    
    long long int result = 0;
    for(int j = 0 ;j < TOTAL_BATTERIES;j++){
      result += batteries[TOTAL_BATTERIES - 1 - j ] * pow(10, j);
    }
    total += result;
  }
  const bool equals = resultContent == (std::to_string(total));
  cout << resultContent << "|" << total << "|" << (equals ? "EQUALS" : "RIP");
  return 0;
}