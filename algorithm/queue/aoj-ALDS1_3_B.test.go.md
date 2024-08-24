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
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
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
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes:
    PROBLEM: https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_B
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/queue/aoj-ALDS1_3_B.test.go\n"
  code: "//go:build ignore\n\n// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_B\n\
    \npackage main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\"\
    strconv\"\n\n\t\"github.com/today2098/algorithm-go/algorithm/queue\"\n)\n\nvar\
    \ sc = bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\nfunc\
    \ getInt() int {\n\tsc.Scan()\n\telem, err := strconv.Atoi(sc.Text())\n\tif err\
    \ != nil {\n\t\tpanic(err)\n\t}\n\treturn elem\n}\n\nfunc getString() string {\n\
    \tsc.Scan()\n\treturn sc.Text()\n}\n\nfunc out(x ...any) {\n\tfmt.Fprintln(wr,\
    \ x...)\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanWords)\n\tsc.Buffer([]byte{},\
    \ math.MaxInt32)\n\tdefer wr.Flush()\n\n\tn, q := getInt(), getInt()\n\n\ttype\
    \ task struct {\n\t\tname string\n\t\ttime int\n\t}\n\tque := queue.NewQueue[*task]()\n\
    \tfor i := 0; i < n; i++ {\n\t\tt := &task{}\n\t\tt.name, t.time = getString(),\
    \ getInt()\n\t\tque.Push(t)\n\t}\n\n\tnow := 0\n\tfor !que.Empty() {\n\t\tt :=\
    \ que.Pop()\n\t\tif t.time <= q {\n\t\t\tnow += t.time\n\t\t\tout(t.name, now)\n\
    \t\t} else {\n\t\t\tnow += q\n\t\t\tt.time -= q\n\t\t\tque.Push(t)\n\t\t}\n\t\
    }\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/stack/stack.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
  isVerificationFile: true
  path: algorithm/queue/aoj-ALDS1_3_B.test.go
  requiredBy:
  - main.go
  - algorithm/stack/stack.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  timestamp: '2024-08-24 17:41:38+09:00'
  verificationStatus: TEST_ACCEPTED
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
documentation_of: algorithm/queue/aoj-ALDS1_3_B.test.go
layout: document
redirect_from:
- /verify/algorithm/queue/aoj-ALDS1_3_B.test.go
- /verify/algorithm/queue/aoj-ALDS1_3_B.test.go.html
title: algorithm/queue/aoj-ALDS1_3_B.test.go
---
