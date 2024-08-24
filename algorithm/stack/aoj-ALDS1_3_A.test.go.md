---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/deque.go
    title: algorithm/deque/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/deque.go
    title: algorithm/deque/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
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
    RuntimeError: bundler is not specified: algorithm/stack/aoj-ALDS1_3_A.test.go\n"
  code: "//go:build ignore\n\n// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A\n\
    \npackage main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\"\
    strconv\"\n\t\"strings\"\n\n\t\"github.com/today2098/algorithm-go/algorithm/stack\"\
    \n)\n\nvar sc = bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\
    \nfunc out(x ...any) {\n\tfmt.Fprintln(wr, x...)\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanLines)\n\
    \tsc.Buffer([]byte{}, math.MaxInt32)\n\tdefer wr.Flush()\n\n\tsc.Scan()\n\tquery\
    \ := strings.Split(sc.Text(), \" \")\n\n\tst := stack.NewStack[int]()\n\tfor _,\
    \ elem := range query {\n\t\tnum, err := strconv.Atoi(elem)\n\t\tif err == nil\
    \ {\n\t\t\tst.Push(num)\n\t\t} else {\n\t\t\tif elem == \"+\" {\n\t\t\t\ttmp :=\
    \ st.Pop() + st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t} else if elem == \"-\" {\n\
    \t\t\t\ttmp := -st.Pop() + st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t} else {\n\t\t\
    \t\ttmp := st.Pop() * st.Pop()\n\t\t\t\tst.Push(tmp)\n\t\t\t}\n\t\t}\n\t}\n\n\t\
    ans := st.Pop()\n\tout(ans)\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/stack.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
  isVerificationFile: true
  path: algorithm/stack/aoj-ALDS1_3_A.test.go
  requiredBy:
  - main.go
  - algorithm/stack/stack.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  timestamp: '2024-08-24 17:41:38+09:00'
  verificationStatus: TEST_ACCEPTED
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
documentation_of: algorithm/stack/aoj-ALDS1_3_A.test.go
layout: document
redirect_from:
- /verify/algorithm/stack/aoj-ALDS1_3_A.test.go
- /verify/algorithm/stack/aoj-ALDS1_3_A.test.go.html
title: algorithm/stack/aoj-ALDS1_3_A.test.go
---
