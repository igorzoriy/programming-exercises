(define (abs x)
    (if (> x 0)
        x
        (- x)))

(define (square x)
    (* x x))

(define (average x y)
    (/ (+ x y) 2))

(define (sqrt-good-enough guess x) (< (abs (- (square guess) x)) 0.0001))

(define (sqrt-improve guess x) (average guess (/ x guess)))

(define (sqrt-iter guess x)
    (if (sqrt-good-enough guess x)
        guess
        (sqrt-iter
            (sqrt-improve guess x)
            x)))

(define (sqrt x)
    (sqrt-iter 1.0 x))

(writeln (sqrt 2))
(writeln (sqrt 3))
(writeln (sqrt 4))
(writeln (sqrt 30625))
