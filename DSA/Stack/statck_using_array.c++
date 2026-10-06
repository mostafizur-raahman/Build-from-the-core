#include <bits/stdc++.h>
using namespace std;

class Stack {
    int top = -1;
    int st[5];

  public:
    void push(int x) {
        if (top == 4) {
            cout << "Stack is full, remove one to push." << endl;
            return;
        }

        top++;
        st[top] = x;
    }

    void pop() {
        if (top == -1) {
            cout << "Stack is empty, you cannot pop." << endl;
            return;
        }

        top--;
    }

    int size() {
        return top + 1;
    }

    int tops() {
        if (top == -1) {
            cout << "Stack is empty." << endl;
            return -1;
        }

        return st[top];
    }

    void prints() {
        for (int i = 0; i <= top; i++) {
            cout << st[i] << " ";
        }
        cout << endl;
    }
};

int main() {
    Stack st;

    st.push(10);
    st.push(20);
    st.push(30);

    st.prints();

    cout << "Top: " << st.tops() << endl;
    cout << "Size: " << st.size() << endl;

    st.pop();

    cout << "Top after pop: " << st.tops() << endl;

    st.prints();
}