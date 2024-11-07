# Graphite
## An AOT Compiler targetting the LLVM IR written in Go.

```
test.gr
```
```c
def add(x, y) {
    let z = x + y
    z = z * z - x + y
    ret z
}

def main() {
    add(5, 3)
    ret 0
}
```
```
compiles to...
```
```
test.ll
```
```c
define double @add(double %x, double %y) {
entry:
        %0 = alloca double
        store double %x, double* %0
        %1 = alloca double
        store double %y, double* %1
        %2 = alloca double
        %3 = load double, double* %0
        %4 = load double, double* %1
        %5 = add double %3, %4
        store double %5, double* %2
        %6 = load double, double* %2
        %7 = load double, double* %2
        %8 = mul double %6, %7
        %9 = load double, double* %0
        %10 = sub double %8, %9
        %11 = load double, double* %1
        %12 = add double %10, %11
        store double %12, double* %2
        %13 = load double, double* %2
        ret double %13
}

define double @main() {
entry:
        %0 = call double @add(double 5.0, double 3.0)
        ret double 0.0
}
```