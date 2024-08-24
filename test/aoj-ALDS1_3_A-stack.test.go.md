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
    path: main.go
    title: main.go
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
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes:
    PROBLEM: https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: test/aoj-ALDS1_3_A-stack.test.go\n"
  code: "//go:build ignore\n\n// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A\n\
    \npackage main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\"\
    strconv\"\n\t\"strings\"\n\n\t\"github.com/today2098/algorithm-go/algorithm\"\n\
    )\n\nvar sc = bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\
    \nfunc out(x ...any) {\n\tfmt.Fprintln(wr, x...)\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanLines)\n\
    \tsc.Buffer([]byte{}, math.MaxInt32)\n\tdefer wr.Flush()\n\n\tsc.Scan()\n\tquery\
    \ := strings.Split(sc.Text(), \" \")\n\n\tst := algorithm.NewStack[int]()\n\t\
    for _, elem := range query {\n\t\tnum, err := strconv.Atoi(elem)\n\t\tif err ==\
    \ nil {\n\t\t\tst.Push(num)\n\t\t} else {\n\t\t\tif elem == \"+\" {\n\t\t\t\t\
    tmp := st.Pop() + st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t} else if elem == \"-\"\
    \ {\n\t\t\t\ttmp := -st.Pop() + st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t} else {\n\
    \t\t\t\ttmp := st.Pop() * st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t}\n\t\t}\n\t}\n\
    \n\tans := st.Pop()\n\tout(ans)\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/deque.go
  - algorithm/queue.go
  - algorithm/stack.go
  isVerificationFile: true
  path: test/aoj-ALDS1_3_A-stack.test.go
  requiredBy:
  - main.go
  - algorithm/deque.go
  - algorithm/queue.go
  - algorithm/stack.go
  timestamp: '2024-08-24 15:16:21+09:00'
  verificationStatus: TEST_ACCEPTED
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
documentation_of: test/aoj-ALDS1_3_A-stack.test.go
layout: document
redirect_from:
- /verify/test/aoj-ALDS1_3_A-stack.test.go
- /verify/test/aoj-ALDS1_3_A-stack.test.go.html
title: test/aoj-ALDS1_3_A-stack.test.go
---
