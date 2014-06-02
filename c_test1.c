#include <stdio.h>

#define MAX 100000

int is_prime(int n) {
    int i;
    for (i = 2; i < n; i++) {
        if (n % i == 0) {
            return 0;
        }
    }
    return 1;
}

int main() {
    int i;
    int count = 0;
    for (i = 2; i <= MAX; i++) {
        if (is_prime(i)) {
            ++count;
        }
    }
    printf("1-%d: %d prime numbers\n", MAX, count);
    return 0;
}