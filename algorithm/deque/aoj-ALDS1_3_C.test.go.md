---
data:
  _extendedDependsOn:
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
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
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
    PROBLEM: https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_C
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/deque/aoj-ALDS1_3_C.test.go\n"
  code: "//go:build ignore\n\n// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_C\n\
    \npackage main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\"\
    strconv\"\n\n\t\"github.com/today2098/algorithm-go/algorithm/deque\"\n)\n\nvar\
    \ sc = bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\nfunc\
    \ getInt() int {\n\tsc.Scan()\n\telem, err := strconv.Atoi(sc.Text())\n\tif err\
    \ != nil {\n\t\tpanic(err)\n\t}\n\treturn elem\n}\n\nfunc getString() string {\n\
    \tsc.Scan()\n\treturn sc.Text()\n}\n\nfunc outArray[T any](arr []T) {\n\tfor i\
    \ := 0; i < len(arr)-1; i++ {\n\t\tfmt.Fprintf(wr, \"%v \", arr[i])\n\t}\n\tif\
    \ len(arr) > 0 {\n\t\tfmt.Fprintf(wr, \"%v\", arr[len(arr)-1])\n\t}\n\tfmt.Fprintf(wr,\
    \ \"\\n\")\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanWords)\n\tsc.Buffer([]byte{},\
    \ math.MaxInt32)\n\tdefer wr.Flush()\n\n\tn := getInt()\n\n\tdq := deque.NewDeque[int]()\n\
    \tfor i := 0; i < n; i++ {\n\t\tcmd := getString()\n\n\t\tif cmd == \"insert\"\
    \ {\n\t\t\tx := getInt()\n\n\t\t\tdq.PushFront(x)\n\t\t} else if cmd == \"delete\"\
    \ {\n\t\t\tx := getInt()\n\n\t\t\tnode := dq.Data.Front()\n\t\t\tfor node != nil\
    \ && node.Value.(int) != x {\n\t\t\t\tnode = node.Next()\n\t\t\t}\n\t\t\tif node\
    \ != nil {\n\t\t\t\tdq.Data.Remove(node)\n\t\t\t}\n\t\t} else if cmd == \"deleteFirst\"\
    \ {\n\t\t\tdq.PopFront()\n\t\t} else {\n\t\t\tdq.PopBack()\n\t\t}\n\t}\n\n\tans\
    \ := []int{}\n\tnode := dq.Data.Front()\n\tfor node != nil {\n\t\tans = append(ans,\
    \ node.Value.(int))\n\t\tnode = node.Next()\n\t}\n\n\toutArray(ans)\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/stack/stack.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  isVerificationFile: true
  path: algorithm/deque/aoj-ALDS1_3_C.test.go
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
  - algorithm/queue/aoj-ALDS1_3_B.test.go
documentation_of: algorithm/deque/aoj-ALDS1_3_C.test.go
layout: document
redirect_from:
- /verify/algorithm/deque/aoj-ALDS1_3_C.test.go
- /verify/algorithm/deque/aoj-ALDS1_3_C.test.go.html
title: algorithm/deque/aoj-ALDS1_3_C.test.go
---
