#include <stdio.h>
#include <unistd.h>

int main() {
    printf("=== Process Information ===\n");
    printf("Process ID (PID): %d\n", getpid());
    printf("Parent Process ID (PPID): %d\n", getppid());
    printf("User ID: %d\n", getuid());
    printf("Group ID: %d\n", getgid());
    return 0;
}
