#include <iostream>
#include <string>
#include <list>
using namespace std;
int FirstTrebuchet(string line){
    int x = -1,y = -1;
    int lineLength = line.length();

    for(int i = 0 ; i < lineLength; i++){
        if(x != -1 && y != -1){
            break;
        }
        if(x == -1 && isdigit(line[i])){
            x = line[i] - '0';
        }
        if(y == -1 && isdigit(line[lineLength - i - 1])){
            y = line[lineLength - i - 1] - '0';
        }
    }
    if(x == -1 && y == -1){
        return -1;
    }
    int sum;
    sum = (x * 10);
    sum += y == -1 ? x : y;
    return sum;
}
int CompareStrings(const string& str, const string& full){
    if(str.length() < full.length()){
        return -1;
    }
    int matches = 0;
    for(int i = 0; i < full.length(); i++){
        if(str[i] == full[i]){
            matches++;
        }
    }
    return matches == full.length() ? 1 : -1;
}
int SecondTrebuchet(string line){


    int x = 0 ,y = 0;
    string words[9] = {"one","two","three","four","five","six","seven","eight","nine"};

    for(int i = 0; i < line.length(); i++){
        char c = line[i];
        if(isdigit(c)){
            if(x == 0){
                x = c - '0';
            }
            y = c - '0';
            continue;
        }

        string currWord = line.substr(i, 5);

        for(int j = 0 ; j < 9; j++){
            int result = CompareStrings(currWord, words[j]);
            if(result == 1){
                if(x == 0){
                    x = j + 1;
                    y = j + 1;
                }else{
                    y = j + 1;
                }
                string temp = words[j];
                break;
            }
        }
    }
    return x * 10 + y;
}
int main() {
    string inputLine;
    int total = 0;
    for(string line; getline( cin, line ); )
    {
        int value = SecondTrebuchet(line);
        //cout << value << endl;
        total += value;
    }
    cout << total << endl;
    return 0;
}
