#include <stdio.h>
#include <unistd.h>

int main() {
    pid_t pid = fork();
    
    if (pid < 0) {
        // Fork failed
        printf("Fork failed!\n");
        return 1;
    }
    else if (pid == 0) {
        // Child process (pid == 0 for child)
        printf("[CHILD] My PID: %d\n", getpid());
        printf("[CHILD] My Parent PID: %d\n", getppid());
    }
    else {
        // Parent process (pid > 0)
        printf("[PARENT] My PID: %d\n", getpid());
        printf("[PARENT] My Child PID: %d\n", pid);
        sleep(1);  // Give child time to execute
    }
    
    return 0;
}
