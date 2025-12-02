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
  int dial = 50; 
  for(auto v : inputLines ){
    const char direction = v[0];
    int intVal = std::stoi(v.substr(1)) ;
    
    while(intVal > 0){
      const int dialBefore = dial;
      if(intVal > 100){        
        total += 1;
        intVal -= 100;
        continue;
      }

      if(direction == 'L'){
        dial -= intVal;
        if(dial < 0){
          if(dialBefore != 0 && dialBefore != 100){
            total += 1;
          }      
          dial = 100 + dial;
        }
      }else if(direction == 'R'){
        dial += intVal;
        if(dial > 100){
          if(dialBefore != 0 && dialBefore != 100){
            total += 1;
          }

          dial = dial - 100;
        }
      }

      intVal -= 100;
      if(dial == 0 || dial == 100){
        total += 1;
      }
    }

  }
  const bool equals = resultContent == (std::to_string(total));
  cout << resultContent << "|" << total << "|" << (equals ? "EQUALS" : "RIP");
  return 0;
}