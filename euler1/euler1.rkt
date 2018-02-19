(define (fsum-new-sum current sum)
    (if (or
            (= 0 (modulo current 3))
            (= 0 (modulo current 5))
        )
            (+ sum current)
            sum
    )
)

(define (fsum-iter current end sum)
    (if (< current end)
        (fsum-iter
            (+ current 1)
            end
            (fsum-new-sum current sum)
        )
        sum
    )
)

(define (fsum start end)
    (fsum-iter start end 0)
)


(writeln (fsum 0 1000))
