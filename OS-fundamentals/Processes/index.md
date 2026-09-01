# OS Fundamentals: Processes & Execution 🖥️

## Table of Contents
1. [What is a Process?](#1-what-is-a-process)
2. [fork(), exec(), and Process Creation](#2-fork-exec-and-process-creation)
3. [Process States & Lifecycle](#3-process-states--lifecycle)
4. [Threads vs Processes](#4-threads-vs-processes)
5. [Signals & Inter-Process Communication](#5-signals--inter-process-communication)

---

## 1. What is a Process?

### Definition
A **process** is a running instance of a program with its own:
- Memory space (stack, heap, data)
- File descriptors
- Process ID (PID)
- Resources (CPU time, I/O)

### Code Example 1: Getting Process Information

**File: `process_info.c`**
```c
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
```

**Compilation:**
```bash
gcc process_info.c -o process_info
```

**Output:**
```
=== Process Information ===
Process ID (PID): 12345
Parent Process ID (PPID): 5432
User ID: 1000
Group ID: 1000
```

---

## 2. fork(), exec(), and Process Creation

### fork() - Create Child Process

**File: `fork_example.c`**
```c
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
```

**Compilation & Execution:**
```bash
gcc fork_example.c -o fork_example
./fork_example
```

**Output:**
```
[PARENT] My PID: 12345
[PARENT] My Child PID: 12346
[CHILD] My PID: 12346
[CHILD] My Parent PID: 12345
```

---

### exec() - Replace Process Image

**File: `exec_example.c`**
```c
#include <stdio.h>
#include <unistd.h>

int main() {
    printf("Before exec - PID: %d\n", getpid());
    printf("Executing 'ls -l' command...\n\n");
    
    // exec replaces current process with 'ls'
    execl("/bin/ls", "ls", "-l", "/tmp", NULL);
    
    // This line will NEVER execute if exec succeeds
    printf("After exec - This won't print\n");
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc exec_example.c -o exec_example
./exec_example
```

**Output:**
```
Before exec - PID: 12347
Executing 'ls -l' command...

total 48
drwxr-xr-x  2 user user  4096 Aug 20 10:15 file1.txt
drwxr-xr-x  2 user user  4096 Aug 20 10:16 file2.txt
-rw-r--r--  1 user user   256 Aug 20 10:17 file3.log
```

---

### fork() + exec() - Process Creation Pattern

**File: `fork_exec_example.c`**
```c
#include <stdio.h>
#include <unistd.h>
#include <sys/wait.h>

int main() {
    pid_t pid = fork();
    
    if (pid == 0) {
        // Child process - execute new program
        printf("[CHILD] PID %d: Executing 'echo' command\n", getpid());
        execl("/bin/echo", "echo", "Hello from child process!", NULL);
    }
    else {
        // Parent process - wait for child
        printf("[PARENT] Created child with PID %d\n", pid);
        wait(NULL);  // Wait for child to finish
        printf("[PARENT] Child process completed\n");
    }
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc fork_exec_example.c -o fork_exec_example
./fork_exec_example
```

**Output:**
```
[PARENT] Created child with PID 12348
[CHILD] PID 12348: Executing 'echo' command
Hello from child process!
[PARENT] Child process completed
```

---

## 3. Process States & Lifecycle

### Process State Diagram

```
┌─────────────────────────────────────────────┐
│       PROCESS STATE TRANSITIONS              │
└─────────────────────────────────────────────┘

    NEW
     ↓ (Admitted)
   READY ←─────────────────────────┐
     ↓ (Dispatch/Scheduled)        │
  RUNNING                           │
     ↓                              │
  WAITING (I/O, Event)              │
     ↓ (I/O Complete)               │
   READY ────────────────────────────┘
     ↓ (Dispatch)
  RUNNING
     ↓ (Exit)
 TERMINATED
```

### Process States Explanation

| State | Description | Example |
|-------|-------------|---------|
| **New** | Process created but not yet admitted to pool | Fork just called |
| **Ready** | Waiting for CPU time | Multiple processes in queue |
| **Running** | Currently executing on CPU | Process actively running |
| **Waiting/Blocked** | Waiting for I/O or event | Reading from disk/network |
| **Terminated** | Process finished or killed | Process completed execution |

### Code Example: Simulating Process States

**File: `process_states.c`**
```c
#include <stdio.h>
#include <unistd.h>
#include <sys/wait.h>

int main() {
    printf("=== Process Lifecycle Demo ===\n\n");
    
    printf("[NEW] Creating process...\n");
    pid_t pid = fork();
    
    if (pid == 0) {
        // Child process
        printf("[READY] Child ready to run\n");
        printf("[RUNNING] Child is running (PID: %d)\n", getpid());
        
        printf("[WAITING] Child waiting for I/O (reading file)...\n");
        FILE *file = fopen("/etc/passwd", "r");
        if (file) {
            char buffer[256];
            fgets(buffer, sizeof(buffer), file);
            printf("[RUNNING] I/O complete, got: %s", buffer);
            fclose(file);
        }
        
        printf("[TERMINATED] Child process exiting\n");
    }
    else {
        // Parent process
        printf("[RUNNING] Parent waiting for child (PID: %d)\n", pid);
        wait(NULL);
        printf("[RUNNING] Parent: Child terminated\n");
    }
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc process_states.c -o process_states
./process_states
```

**Output:**
```
=== Process Lifecycle Demo ===

[NEW] Creating process...
[READY] Child ready to run
[RUNNING] Child is running (PID: 12349)
[WAITING] Child waiting for I/O (reading file)...
[RUNNING] I/O complete, got: root:x:0:0:root:/root:/bin/bash
[TERMINATED] Child process exiting
[RUNNING] Parent: Child terminated
```

---

## 4. Threads vs Processes

### Comparison Table

| Feature | **Process** | **Thread** |
|---------|-----------|----------|
| **Memory** | Separate memory space | Shared memory space |
| **Creation Time** | Slower (more overhead) | Faster (less overhead) |
| **Context Switch** | Slower, more expensive | Faster, lighter |
| **Isolation** | Highly isolated | Minimal isolation |
| **IPC** | Complex (pipes, sockets) | Simple (shared variables) |
| **Safety** | Safe by default | Race conditions possible |
| **Resource Usage** | High | Low |

---

### Multi-Process Example

**File: `multiprocess.c`**
```c
#include <stdio.h>
#include <unistd.h>
#include <sys/wait.h>

int main() {
    printf("=== Multi-Process Example ===\n");
    printf("Parent PID: %d\n\n", getpid());
    
    // Create 3 child processes
    for (int i = 0; i < 3; i++) {
        pid_t pid = fork();
        
        if (pid == 0) {
            // Child process
            printf("Child %d - PID: %d, Parent: %d\n", i+1, getpid(), getppid());
            sleep(1);  // Simulate work
            return 0;
        }
    }
    
    // Parent waits for all children
    printf("Parent: Waiting for children...\n");
    for (int i = 0; i < 3; i++) {
        wait(NULL);
    }
    printf("Parent: All children completed\n");
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc multiprocess.c -o multiprocess
./multiprocess
```

**Output:**
```
=== Multi-Process Example ===
Parent PID: 12350

Child 1 - PID: 12351, Parent: 12350
Child 2 - PID: 12352, Parent: 12350
Child 3 - PID: 12353, Parent: 12350
Parent: Waiting for children...
Parent: All children completed
```

---

### Multi-Threading Example

**File: `multithreading.c`**
```c
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

int shared_counter = 0;  // Shared memory between threads

void* thread_function(void* arg) {
    int thread_id = *(int*)arg;
    
    printf("Thread %d - Starting (Thread ID: %lu)\n", thread_id, pthread_self());
    
    for (int i = 0; i < 3; i++) {
        shared_counter++;
        printf("Thread %d - Counter: %d\n", thread_id, shared_counter);
        sleep(1);
    }
    
    printf("Thread %d - Ending\n", thread_id);
    return NULL;
}

int main() {
    printf("=== Multi-Threading Example ===\n");
    printf("Main Thread PID: %d\n\n", getpid());
    
    pthread_t threads[2];
    int thread_ids[2] = {1, 2};
    
    // Create 2 threads
    for (int i = 0; i < 2; i++) {
        pthread_create(&threads[i], NULL, thread_function, &thread_ids[i]);
    }
    
    printf("Main: Waiting for threads...\n\n");
    
    // Wait for threads to complete
    for (int i = 0; i < 2; i++) {
        pthread_join(threads[i], NULL);
    }
    
    printf("\nMain: All threads completed\n");
    printf("Final shared_counter: %d\n", shared_counter);
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc -pthread multithreading.c -o multithreading
./multithreading
```

**Output:**
```
=== Multi-Threading Example ===
Main Thread PID: 12354

Main: Waiting for threads...

Thread 1 - Starting (Thread ID: 139876543210752)
Thread 2 - Starting (Thread ID: 139876543198208)
Thread 1 - Counter: 1
Thread 2 - Counter: 2
Thread 1 - Counter: 3
Thread 2 - Counter: 4
Thread 1 - Counter: 5
Thread 2 - Counter: 6
Thread 1 - Ending
Thread 2 - Ending

Main: All threads completed
Final shared_counter: 6
```

---

## 5. Signals & Inter-Process Communication

### What are Signals?

**Signals** are software interrupts that provide a way to handle asynchronous events.

### Common Signals

| Signal | Number | Meaning |
|--------|--------|---------|
| **SIGHUP** | 1 | Hangup |
| **SIGINT** | 2 | Interrupt (Ctrl+C) |
| **SIGQUIT** | 3 | Quit |
| **SIGKILL** | 9 | Kill (cannot be caught) |
| **SIGTERM** | 15 | Termination signal |
| **SIGSTOP** | 19 | Stop (cannot be caught) |
| **SIGCHLD** | 17 | Child process terminated |
| **SIGUSR1** | 10 | User-defined 1 |
| **SIGUSR2** | 12 | User-defined 2 |

---

### Code Example 1: Signal Handling

**File: `signal_handler.c`**
```c
#include <stdio.h>
#include <signal.h>
#include <unistd.h>
#include <stdlib.h>

void signal_handler(int sig) {
    if (sig == SIGINT) {
        printf("\n[CAUGHT] SIGINT received (Ctrl+C)\n");
        printf("Performing cleanup...\n");
        exit(0);
    }
    else if (sig == SIGUSR1) {
        printf("[CAUGHT] SIGUSR1 received\n");
    }
}

int main() {
    printf("=== Signal Handling Example ===\n");
    printf("Press Ctrl+C to send SIGINT\n");
    printf("Or run: kill -USR1 %d\n\n", getpid());
    
    // Register signal handlers
    signal(SIGINT, signal_handler);   // Ctrl+C
    signal(SIGUSR1, signal_handler);  // Custom signal
    
    int counter = 0;
    while (1) {
        printf("Counter: %d (Process running...)\n", ++counter);
        sleep(1);
    }
    
    return 0;
}
```

**Compilation:**
```bash
gcc signal_handler.c -o signal_handler
```

**Execution & Output:**
```bash
./signal_handler
```

**Output (with Ctrl+C after 3 iterations):**
```
=== Signal Handling Example ===
Press Ctrl+C to send SIGINT
Or run: kill -USR1 12355

Counter: 1 (Process running...)
Counter: 2 (Process running...)
Counter: 3 (Process running...)
^C
[CAUGHT] SIGINT received (Ctrl+C)
Performing cleanup...
```

---

### Code Example 2: Parent-Child Signal Communication

**File: `signal_ipc.c`**
```c
#include <stdio.h>
#include <signal.h>
#include <unistd.h>
#include <sys/wait.h>

void child_handler(int sig) {
    printf("[PARENT] Caught SIGUSR1 from child!\n");
}

int main() {
    printf("=== Parent-Child Signal Communication ===\n");
    
    pid_t pid = fork();
    
    if (pid == 0) {
        // Child process
        printf("[CHILD] Child PID: %d\n", getpid());
        printf("[CHILD] Sleeping for 2 seconds...\n");
        sleep(2);
        
        printf("[CHILD] Sending SIGUSR1 to parent\n");
        kill(getppid(), SIGUSR1);
        
        printf("[CHILD] Exiting\n");
    }
    else {
        // Parent process
        printf("[PARENT] Parent PID: %d\n", getpid());
        printf("[PARENT] Created child: %d\n\n", pid);
        
        signal(SIGUSR1, child_handler);
        
        wait(NULL);
        printf("[PARENT] Child process terminated\n");
    }
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc signal_ipc.c -o signal_ipc
./signal_ipc
```

**Output:**
```
=== Parent-Child Signal Communication ===
[PARENT] Parent PID: 12356
[PARENT] Created child: 12357

[CHILD] Child PID: 12357
[CHILD] Sleeping for 2 seconds...
[CHILD] Sending SIGUSR1 to parent
[CHILD] Exiting
[PARENT] Caught SIGUSR1 from child!
[PARENT] Child process terminated
```

---

### Code Example 3: Inter-Process Communication (IPC) - Pipes

**File: `pipe_ipc.c`**
```c
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <sys/wait.h>

int main() {
    printf("=== IPC using Pipes ===\n");
    
    int pipefd[2];  // pipefd[0] = read, pipefd[1] = write
    pipe(pipefd);
    
    pid_t pid = fork();
    
    if (pid == 0) {
        // Child process - Writer
        close(pipefd[0]);  // Close read end
        
        char* message = "Hello from Child!";
        printf("[CHILD] Sending message: %s\n", message);
        write(pipefd[1], message, strlen(message) + 1);
        close(pipefd[1]);
    }
    else {
        // Parent process - Reader
        close(pipefd[1]);  // Close write end
        
        char buffer[100];
        read(pipefd[0], buffer, sizeof(buffer));
        printf("[PARENT] Received message: %s\n", buffer);
        close(pipefd[0]);
        
        wait(NULL);
    }
    
    return 0;
}
```

**Compilation & Execution:**
```bash
gcc pipe_ipc.c -o pipe_ipc
./pipe_ipc
```

**Output:**
```
=== IPC using Pipes ===
[CHILD] Sending message: Hello from Child!
[PARENT] Received message: Hello from Child!
```

---

## Quick Reference Guide

### Commands
```bash
# View running processes
ps aux

# Check specific process
ps -p <PID>

# Send signal to process
kill -SIGNAL <PID>
kill -9 <PID>           # SIGKILL
kill -15 <PID>          # SIGTERM

# View process tree
pstree

# Monitor processes
top
htop

# Create background process
./program &

# Kill background process
fg              # Bring to foreground
Ctrl+C          # Kill it
```

### Function Reference

```c
// Process Management
pid_t fork();                           // Create child process
int execl(const char *path, ...);      // Execute program
int wait(int *status);                  // Wait for child
pid_t getpid();                         // Get current PID
pid_t getppid();                        // Get parent PID

// Signals
int signal(int sig, void (*func)(int));  // Register handler
int kill(pid_t pid, int sig);           // Send signal

// Threads
int pthread_create(...);                // Create thread
int pthread_join(...);                  // Wait for thread
pthread_t pthread_self();               // Get thread ID

// IPC
int pipe(int pipefd[2]);                // Create pipe
ssize_t read(int fd, void *buf, size_t count);
ssize_t write(int fd, const void *buf, size_t count);
```

---

## Summary

✅ **Processes** are independent programs with separate memory
✅ **fork()** creates child processes, **exec()** replaces program image  
✅ **Process States** cycle through Ready → Running → Waiting → Terminated  
✅ **Threads** share memory and are faster but less isolated  
✅ **Signals** enable inter-process communication and event handling  

---
