# go-history
Command line utility for working with bash history, I got frustrated trying to do this effieciently in bash: 

```
2029  history
2030  history | grep export | grep -v '^ *[0-9]* *history' | uniq -u | awk '{$1=""; print $0}' | wc -l
2031  history | grep export | grep -v '^ *[0-9]* *history' | sort | uniq -u | awk '{$1=""; print $0}' | wc -l
2032  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u | wc -l
2033  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u
2034  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u 
2035  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u | wc -l
2036  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u | >> build.txt
2037  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u | > build.txt
2038  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u | wc -l
2039  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u 
2040  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u > test
2041  cat test
2042  history | grep export | grep -v '^ *[0-9]* *history' | awk '{$1=""; print $0}' | sort | uniq -u > test
2043  cat test
2044  history | grep export | grep -v '^ *[0-9]* *history' | grep -v '^ *history' | awk '{$1=""; print $0}' | sort | uniq -u > test
2045  cat test
2046  grep 'export' ~/.bash_history | grep -v '^ *[0-9]* *history' | sort | uniq -u > test
2047  grep 'export*' ~/.bash_history | grep -v '^ *[0-9]* *history' | sort | uniq -u > test
```
