#include <iostream>
#include <map>
#include <utility>
using namespace std;

//Hand Cards string of 5 length
// determine the strength of the hand
// 1. High Card
// 2. Pair
// 3. Two Pair
// 4. Three of a Kind
// 5. Full house
// 6. Four of a Kind
// 7. Five of a kind
map<char, int> cardStrength = {
        {'J', 1},
        {'2', 2},
        {'3', 3},
        {'4', 4},
        {'5', 5},
        {'6', 6},
        {'7', 7},
        {'8', 8},
        {'9', 9},
        {'T', 10},
        {'J', 11},
        {'Q', 12},
        {'K', 13},
        {'A', 14}
};
class PokerHand {
    public:
        PokerHand(){
            hand = "";
            bet = 0;
            strength = 0;
            hiddenHand = "";
        }
        string getHand(){
            return hand;
        }
        int getStrength() const{
            return strength;
        }
        int getBet() const{
            return bet;
        }
        string getHiddenHand(){
            return hiddenHand;
        }
        void setBet(int b){
            this->bet = b;
        }
        void setHand(string h){
            //sort h
            this->hand = std::move(h);
            for(int i = 0; i < 5; i++){
                for(int j = i+1; j < 5; j++){
                    if(cardStrength[h[i]] > cardStrength[h[j]]){
                        char temp = h[i];
                        h[i] = h[j];
                        h[j] = temp;
                    }
                }
            }
            this->hiddenHand = std::move(h);
        }
        void setStrength(){
            map<char,int> m;
            int strongest = 0;
            char strongestKey = -1;
            m['J'] = 0;
            for(char i : hand){
                m[i]++;
                if(i != 'J' && m[i] > strongest){
                    strongest = m[i];
                    strongestKey = i;
                }
            }
            m[strongestKey] += m['J'];
            m['J'] = 0;
            int handStrength = 0;
            for(auto i : m){
                cout << i.first << " " << i.second << endl;
                if(i.second == 2){
                    if(handStrength == 0){
                        handStrength = 1;
                    }else if(handStrength == 1){
                        handStrength = 2;
                    }else if(handStrength == 3){
                        handStrength = 4;
                    }
                }
                if(i.second == 3){
                    if(handStrength == 0){
                        handStrength = 3;
                    }else{
                        handStrength = 4;
                    }
                }
                if(i.second == 4){
                    handStrength = 5;
                }
                if(i.second == 5){
                    handStrength = 6;
                }
            }
            this->strength = handStrength;
        }
        bool isHandStronger(PokerHand p){
            if(strength > p.getStrength()){
                return true;
            }
            if(strength < p.getStrength()){
                return false;
            }
            string myHand = getHand();
            // loop through hand of p and compare with this -> hand if letter of this -> hand > p return true
            for (int i = 0; i < 5; i++){
                if(cardStrength[myHand[i]] > cardStrength[p.getHand()[i]]){
                    return true;
                }else if (cardStrength[myHand[i]] < cardStrength[p.getHand()[i]]){
                    return false;
                }
            }
            return false;
        }

    private:
        string hand;
        string hiddenHand;
        int bet;
        int strength;
};
int main() {
    //getline
    string hand;
    int bet;
    //array of hands
    map<int,PokerHand> hands;
    //hands[0] = PokerHand();
    //hands[0].setBet(10);
    //hands[0].setHand("9T7JJ");
    //hands[0].setStrength();
    //cout << hands[0].getStrength() << endl;
    //return 0;
    int n = 0;
    while(cin >> hand >> bet){
        hands[n].setBet(bet);
        hands[n].setHand(hand);
        hands[n].setStrength();
        n++;
    }
    //sort hands

    //sort hands
    for(int i = 0; i < n; i++){
        for(int j = i+1; j < n; j++){
            if(hands[i].isHandStronger(hands[j])){
                PokerHand temp = hands[i];
                hands[i] = hands[j];
                hands[j] = temp;
            }
        }
    }
    //print hands
    long int sum = 0;
    for(int i = 0; i < n; i++){
        sum += hands[i].getBet() * (i+1);
        cout << i << " " << hands[i].getHand() << " " <<  hands[i].getHiddenHand() << " " << hands[i].getBet() << " " << hands[i].getStrength() << endl;
    }
    cout << sum << endl;
    return 0;
}
