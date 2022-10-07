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

// average of n numbers
null p3() {

    i32 n, a;
    f64 sum attr 0.0;

    print "n = ";
    scanner shiftR n;

    print "Input ";
    print n;
    print " numbers: ";

    i64 i attr 0;

    loop (i lt n) {

        scanner shiftR a;
        sum add a;

        i add 1;
    }

    print "Average: ";
    print sum / n;
}

i32 main() {

    p3 call;

    give 0;
}


