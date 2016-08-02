(define (fib-new-sum current sum)
    (if (= 0 (modulo current 2))
        (+ sum current)
        sum
    )
)

(define (fib-sum x y max sum)
    (if (> y max)
        sum
        (fib-sum y (+ x y) max (fib-new-sum (+ x y) sum))
    )
)

(writeln (fib-sum 1 2 4000000 2))
