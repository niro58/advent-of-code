#include <iostream>
#include <vector>
#include <sstream>
using namespace std;
class Race{
public:
    int CalculateWaysToBeat(){
        long int ways = 0;
        for(long int i = time / 2; i >= 0; i--){
            long int distance = (time - i) * i;
            if(distance > record){
                if(ways == 0 && time % 2 == 0){
                    ways += 1;

                    continue;
                }
                ways += 2;
            }
        }
        cout << endl;
        return ways;
    }
    void SetTime(long int t){
        this->time = t;
    }
    void SetRecord(long int r){
        this->record = r;
    }
private:
    long int time;
    long int record;
};
int main() {
    //check input word by word
    string recordLine;
    string timeLine;

    getline(cin, recordLine);
    getline(cin, timeLine);

    istringstream iss(recordLine);
    istringstream iss2(timeLine);

    long int sum = 1;
    long int calc;

    //throw away first word
    string word;
    iss >> word;
    iss2 >> word;

    long int record;
    long int time;
    for(int i = 0; i < 1;i++){

        iss >> time;
        iss2 >> record;

        cout << "Time: " << time << endl;
        cout << "Record: " << record << endl;

        Race race = {};
        race.SetRecord(record);
        race.SetTime(time);

        int variations = race.CalculateWaysToBeat();
        cout << variations << endl;
        cout << endl;
        sum *= variations;
    }
    cout << sum << endl;
    return 0;
}
