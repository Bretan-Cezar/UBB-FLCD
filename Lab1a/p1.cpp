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

// max of 3 numbers
null p1() {

    f64 n1, n2, n3;

    scanner shiftR n1 shiftR n2 shiftR n3;

    jmp (n1 gte n2 AND n1 gte n3) {

        print "Maximum number is: ";
        print n1;
        print newln;
    }

    jmp (n2 gte n1 AND n2 gte n3) {

        print "Maximum number is: ";
        print n2;
        print newln;
    }

    jmp (n3 gte n1 AND n3 gte n2) {

        print "Maximum number is: ";
        print n3;
        print newln;
    }
}

i32 main() {

    p1 call;

    give 0;
}
