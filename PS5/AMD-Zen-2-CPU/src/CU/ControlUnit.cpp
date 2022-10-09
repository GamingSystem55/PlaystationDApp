
#include <iostream>
using namespace std;

class ControlUnit
{
	public:
		int x;
		string String;
};
 
int main()
{
	ControlUnit cu;
	
	cu.x = 0;
	cu.String="ControlUnit";
	
	cout << cu.x << "\n";
	return 0;
}
