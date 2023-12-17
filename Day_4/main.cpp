#include <iostream>
#include <map>
#include <sstream>
#include <vector>
using namespace std;
void firstV(){
    //string line;
    //int sum = 0;
    //int index = 0;
    //map<int, int> timesToPlay;
    //while (getline(cin, line)) {
    //    string allNumbers = line.substr(line.find(':') + 1, line.length());
    //    string winningNumbers = allNumbers.substr(0, allNumbers.find('|') - 1);
    //    string matchingNumbers = allNumbers.substr(allNumbers.find('|') + 1, allNumbers.length());
//
    //    stringstream ss;
    //    list<int> winningNumbersList;
//
    //    ss << winningNumbers;
    //    string temp;
    //    while (ss >> temp) {
    //        winningNumbersList.push_back(stoi(temp));
    //    }
//
    //    ss.clear();
    //    ss << matchingNumbers;
    //    int matches = 0;
    //    list<int> matchingNumbersList;
    //    while (ss >> temp) {
    //        matchingNumbersList.push_back(stoi(temp));
//
    //    }
    //}
}

int main() {
    map<int,int> timesToPlay;
    timesToPlay.insert(pair<int,int>(0,1));
    string line;
    int index = 0;
    int sum = 0;
    while(getline(cin,line)) {
        string allNumbersRaw = line.substr(line.find(':') + 1, line.length());
        string winningNumbersRaw = allNumbersRaw.substr(0, allNumbersRaw.find('|') - 1);
        string matchingNumbersRaw = allNumbersRaw.substr(allNumbersRaw.find('|') + 1, allNumbersRaw.length());

        stringstream ss;

        vector<int> winningNumbersList;
        ss << winningNumbersRaw;
        string temp;
        while (ss >> temp) {
            winningNumbersList.push_back(stoi(temp));
        }
        ss.clear();
        ss << matchingNumbersRaw;
        int matches = 0;
        cout << "__________________________" << timesToPlay[index] << endl;
        while (ss >> temp) {
            for (int i : winningNumbersList) {
                if (stoi(temp) == i) {
                    matches++;
                }
            }
        }
        int playedTimes = timesToPlay[index];
        sum += timesToPlay[index];

        timesToPlay.erase(timesToPlay[index]);


        cout << "TOTAL MATCHES: " << matches << " Played times : " << playedTimes << endl;
        cout << "Before " << timesToPlay[index + 1] << endl;

        //first element += 1;
        timesToPlay[index + 1] += 1;
        for(int i = 0; i < matches; i++){
            timesToPlay[index + i] += 1 * playedTimes;
        }
        cout << "After " << timesToPlay[index + 1] << endl;
        index++;
        //print timestoplay map
        for(auto i : timesToPlay){
            cout << i.first << " " << i.second << endl;
        }
    }
    cout << sum << endl;
    return 0;
}