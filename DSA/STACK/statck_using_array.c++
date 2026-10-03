#include <bits/stdc++.h>
using namespace std;
class myStack {
    int *arr;
    int capacity;
    int top;

  public:
    myStack(int cap) {
        capacity = cap;
        arr = new int[capacity];
        top = -1;
    }
    void push(int val) {
        if (top == capacity - 1) {
            cout << "Stack overflow..." << endl;
            return;
        }
        cout << val << " is push into array" << endl;
        arr[++top] = val;
        printArr();
    }
    void pop() {
        if (top == -1) {
            cout << "Stack underflow..." << endl;
            return;
        }
        int val = arr[top--];
        cout << val << " is removed from array " << endl;
        printArr();
    }

    int peek() {
        if (top == -1) {
            cout << "Stack is empty..." << endl;
            return -1;
        }
        return arr[top];
    }

    bool isEmpty() {
        return top == -1;
    }
    void printArr() {
        if (top == -1) {
            cout << "Stack is empty..." << endl;
            return;
        }
        cout << "Stack array is " << endl;
        for (int i = top; i >= 0; i--) {
            cout << arr[i] << " ";
        }
        cout << endl;
        cout << "Original array is : ";
        for (int i = 0; i < 5; i++) {
            cout << arr[i] << " ";
        }
        cout << endl;
    }
};

int main(void) {
    myStack s(5);
    s.push(1); // 1
    s.push(2); // 1 2
    s.push(3); // 1 2 3
    s.pop();   // 1 2

    s.printArr(); // 2 1, original array is 1 2 3
    s.push(10);   // 1 2 10
    s.printArr(); // 10 2 1,

    return 0;
}