//
// Created by Cezar Bretan on 2022-10-07.
//
#include <iostream>
#define i32 int
#define f64 double
#define scanner std::cin
#define print std::cout <<
#define newln std::endl
#define jmp if
#define attr =
#define gt >
#define lt <
#define eq ==
#define gte >=
#define lte <=
#define OR ||
#define AND &&
#define shiftL <<
#define shiftR >>
#define null void
#define call ()
#define loop while
#define b bool
#define i64 long long
#define MOD %
#define give return
#define add +=

// prime number
null p2() {

    i64 n, d;
    b result;

    scanner shiftR n;

    jmp (n eq 0 OR n eq 1) {

        print "The given number is not prime.\n";
        give;
    }

    jmp (n eq 2) {

        print "The given number is prime.\n";
        give;
    }

    result attr true;

    jmp (n MOD 2 eq 0) {

        result attr false;
    }

    d attr 3;

    loop (d lte n/2) {

        jmp (n MOD d eq 0) {

            result attr false;
        }

        d add 2;
    }

    jmp (result eq true) {

        print "The given number is prime.\n";
    }
    else {

        print "The given number is not prime.\n";
    }
}

i32 main() {

    p2 call;

    give 0;
}
