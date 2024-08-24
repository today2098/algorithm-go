---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_A-stack.test.go
    title: test/aoj-ALDS1_3_A-stack.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_A-stack.test.go
    title: test/aoj-ALDS1_3_A-stack.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: main.go\n"
  code: "package main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\
    \"strconv\"\n)\n\nvar sc = bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\
    \nfunc getInt() int {\n\tsc.Scan()\n\telem, err := strconv.Atoi(sc.Text())\n\t\
    if err != nil {\n\t\tpanic(err)\n\t}\n\treturn elem\n}\n\nfunc getFloat64() float64\
    \ {\n\tsc.Scan()\n\telem, err := strconv.ParseFloat(sc.Text(), 64)\n\tif err !=\
    \ nil {\n\t\tpanic(err)\n\t}\n\treturn elem\n}\n\nfunc getString() string {\n\t\
    sc.Scan()\n\treturn sc.Text()\n}\n\nfunc getInts(n int) []int {\n\tv := make([]int,\
    \ n)\n\tfor i := 0; i < n; i++ {\n\t\tv[i] = getInt()\n\t}\n\treturn v\n}\n\n\
    func out(x ...any) {\n\tfmt.Fprintln(wr, x...)\n}\n\nfunc outArray[T any](arr\
    \ []T) {\n\tfor i := 0; i < len(arr)-1; i++ {\n\t\tfmt.Fprintf(wr, \"%v \", arr[i])\n\
    \t}\n\tif len(arr) > 0 {\n\t\tfmt.Fprintf(wr, \"%v\", arr[len(arr)-1])\n\t}\n\t\
    fmt.Fprintf(wr, \"\\n\")\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanWords)\n\t\
    sc.Buffer([]byte{}, math.MaxInt32)\n\tdefer wr.Flush()\n\n\tn, m, s := getInt(),\
    \ getFloat64(), getString()\n\tout(n, m, s)\n\n\tv := getInts(n)\n\toutArray(v)\n\
    }\n"
  dependsOn:
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
  - algorithm/deque.go
  - algorithm/queue.go
  - algorithm/stack.go
  isVerificationFile: false
  path: main.go
  requiredBy:
  - algorithm/deque.go
  - algorithm/queue.go
  - algorithm/stack.go
  timestamp: '2024-08-24 15:16:21+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
documentation_of: main.go
layout: document
redirect_from:
- /library/main.go
- /library/main.go.html
title: main.go
---
