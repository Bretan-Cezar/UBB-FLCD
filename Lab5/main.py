from grammar import Grammar
from Parser import ParserRecursiveDescendent

# pr=ParserRecursiveDescendent("Lab5/g1.txt","seq.txt","Lab5/out1.txt")
pr=ParserRecursiveDescendent("Lab5/g2.txt","pif.out","Lab5/out2.txt")
pr.run(pr.sequence)

print(str(pr))