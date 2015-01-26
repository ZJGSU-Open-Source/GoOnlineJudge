package vjudger

import (
	"log"
	"testing"
	// "time"
)

func Test_HDU(t *testing.T) {
	u := &User{Vid: 1000, Lang: 0}
	u.Code = `
#include<iostream>
 
using namespace std;
 
int main(){
   int a,b;
   while(cin>>a>>b){
      cout<<a+b<<endl;
   }
   return 0;
}
	`
	h := &HDUJudger{}
	err := h.Run(u)
	if err != nil {
		t.Error(err)
	}
	log.Println(*u)
}
